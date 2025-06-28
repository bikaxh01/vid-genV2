package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"os/exec"
	"path/filepath"

	"github.com/bikaxh/vid-gen/generator/pkg/prompts"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type SceneGeneration struct {
	ClassName   string `json:"className"`
	Code        string `json:"code"`
	Description string `json:"description"`
	SceneTitle  string `json:"sceneTitle"`
}

func GenerateCode(sceneMetadata prompts.Scene, scenes prompts.Scenes) (*SceneGeneration, error) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GoDotEnvVariable("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type:        genai.TypeObject,
			Description: "Detail of scene",
			Properties: map[string]*genai.Schema{
				"sceneTitle": {
					Type:        genai.TypeString,
					Description: "The title of the scene, describing the main focus or idea.",
				},
				"description": {
					Type:        genai.TypeString,
					Description: "A detailed explanation of what happens in the scene.",
				},
				"className": {
					Type:        genai.TypeString,
					Description: "Name of Class that has been created",
				},
				"code": {
					Type:        genai.TypeString,
					Description: "scene main code of scene",
				},
			},
			Required: []string{"sceneTitle", "description", "code", "className"},
		},
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompts.GetSceneGenerationPrompt(sceneMetadata, scenes)),
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	res := []byte(result.Text())

	var data SceneGeneration
	err = json.Unmarshal(res, &data)

	if err != nil {
		fmt.Println("🔴", err)
		return nil, err
	}

	return &data, nil

}

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("../../.env")

	if err != nil {

		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}

func WriteToFile(fileName, content string) bool {

	mediaPath := filepath.Join("..", "..", "final", "code", fileName+".py")
	absolutePath, err := filepath.Abs(mediaPath)

	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return false
	}

	err = os.WriteFile(absolutePath, []byte(content), 0644)

	if err != nil {
		return false
	}

	return true
}

func CompileFile(fileName string) (bool, error) {

	pyFilePath := filepath.Join("..", "..", "final", "code", fileName+".py")
	pyAbsolutePath, err := filepath.Abs(pyFilePath)

	if err != nil {
		fmt.Println("Error getting absolute path for Python file:", err)
		return false, nil
	}

	mediaDirPath := filepath.Join("..", "..", "final", "media")
	mediaAbsolutePath, err := filepath.Abs(mediaDirPath)
	if err != nil {
		fmt.Println("Error getting absolute path for media dir:", err)
		return false, nil
	}

	cmd := exec.Command("manim", "-qh", pyAbsolutePath, fileName, "--media_dir", mediaAbsolutePath)
	_, err = cmd.CombinedOutput()

	if err != nil {

		return false, err

	}

	return true, nil
}

func FixCode(err string, generatedScene SceneGeneration) {

	history := []*genai.Content{
		genai.NewContentFromText(prompts.GetFixCodePrompt(), genai.RoleModel),
	}
	initialCode, _ := ReadFile(generatedScene.ClassName)
	fmt.Println("Fixing for 🟢", initialCode)

	compilationError := err

	currentPrompt := fmt.Sprintf("Compilation err: ###\n %v \n### current code : ###\n %v \n###", compilationError, initialCode)

	for i := range 5 {
		fmt.Println("Fixing for 🟢", i)
		fmt.Println("histtory for 🟢", generatedScene.SceneTitle,"is",history)
		//pass to llm
		fixedCode := FixCodeLLM(history, currentPrompt)

		//save code to file
		WriteToFile(generatedScene.ClassName, fixedCode)
		//compile
		success, cErr := CompileFile(generatedScene.ClassName)
		if cErr != nil {

			compilationError = fmt.Sprintf("%+v", cErr)
		}
		initialCode, _ = ReadFile(generatedScene.ClassName)
		if success == true {
			break
		}
		history = append(history, genai.NewContentFromText(*initialCode, genai.RoleModel))
	}

}

func ReadFile(fileName string) (*string, error) {
	pyFilePath := filepath.Join("..", "..", "final", "code", fileName+".py")
	pyAbsolutePath, err := filepath.Abs(pyFilePath)

	if err != nil {
		fmt.Println("Error getting absolute path for Python file:", err)
		return nil, err
	}

	d, err := os.ReadFile(pyAbsolutePath)
	data := string(d)
	return &data, nil
}

func FixCodeLLM(history []*genai.Content, currentPrompt string) string {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: GoDotEnvVariable("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type:        genai.TypeObject,
			Description: "Fixed code of scene",
			Properties: map[string]*genai.Schema{
				"code": {
					Type:        genai.TypeString,
					Description: "Fixed code of scene",
				},
			},
			Required: []string{"code"},
		},
	}

	chat, _ := client.Chats.Create(ctx, "gemini-2.5-flash", config, history)
	res, _ := chat.SendMessage(ctx, genai.Part{Text: currentPrompt})

	result := res.Candidates[0].Content.Parts[0].Text

	bytTxt := []byte(result)

	var data map[string]string
	err = json.Unmarshal(bytTxt, &data)

	if err != nil {
		fmt.Println("🔴", err)
		return err.Error()
	}
	return data["code"]
}
