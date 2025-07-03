package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/bikaxh/vid-gen/generator/pkg/prompts"
	"github.com/joho/godotenv"
	

	"os/exec"
	"path/filepath"
)

type SceneGeneration struct {
	ClassName   string `json:"className"`
	Code        string `json:"code"`
	Description string `json:"description"`
	SceneTitle  string `json:"sceneTitle"`
}

func GenerateCode(sceneMetadata prompts.Scene, scenes prompts.Scenes, previousCode string) (*SceneGeneration, error) {

	client := anthropic.NewClient(
		option.WithAPIKey(GoDotEnvVariable("ANTHROPIC_API_KEY")),
	)
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 3024,
		System: []anthropic.TextBlockParam{
			{Text: prompts.GetSceneGenerationSystemPrompt(scenes)},
		},
		Messages: []anthropic.MessageParam{{
			Content: []anthropic.ContentBlockParamUnion{{
				OfText: &anthropic.TextBlockParam{Text: prompts.GetSceneGenerationUserPrompt(previousCode, sceneMetadata)},
			}},
			Role: anthropic.MessageParamRoleUser,
		}},
		Model: anthropic.ModelClaudeOpus4_0,
	})
	if err != nil {
		panic(err.Error())
	}
	
	result := message.Content[0].Text
   
	res := []byte(result)

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

	history := []anthropic.MessageParam{}

	initialCode, _ := ReadFile(generatedScene.ClassName)

	compilationError := err

	currentPrompt := fmt.Sprintf("Compilation err: ###\n %v \n### current code : ###\n %v \n###", compilationError, *initialCode)

	for range 5 {
		fmt.Println("Fixing", generatedScene.SceneTitle)
		history = append(history, anthropic.MessageParam{Content: []anthropic.ContentBlockParamUnion{{OfText: &anthropic.TextBlockParam{
			Text: currentPrompt,
		}}},
			Role: anthropic.MessageParamRoleUser,
		})
		//pass to llm
		fixedCode := FixCodeLLM(history)

		//save code to file
		WriteToFile(generatedScene.ClassName, fixedCode)
		//compile
		success, cErr := CompileFile(generatedScene.ClassName)
		if cErr != nil {

			compilationError = fmt.Sprintf("%+v", cErr)
		}
		initialCode, _ = ReadFile(generatedScene.ClassName)

		if success == true {
		fmt.Printf("History: %+v\n", history)
			fmt.Println("successfully fixed 🟢", generatedScene.SceneTitle)
			return
		}
		history = append(history, anthropic.MessageParam{Content: []anthropic.ContentBlockParamUnion{{OfText: &anthropic.TextBlockParam{
			Text: *initialCode,
		}}},
			Role: anthropic.MessageParamRoleAssistant,
		})
	}

	fmt.Println("Unable to fix 🔴", generatedScene.SceneTitle)

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

func FixCodeLLM(history []anthropic.MessageParam) string {

	client := anthropic.NewClient(
		option.WithAPIKey(GoDotEnvVariable("ANTHROPIC_API_KEY")),
	)
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 3024,
		System: []anthropic.TextBlockParam{
			{Text: prompts.GetFixCodePrompt()},
		},
		Messages: history,
		Model:    anthropic.ModelClaudeOpus4_0,
	})
	if err != nil {
		panic(err.Error())
	}
	result := message.Content[0].Text

	bytTxt := []byte(result)

	var data map[string]string
	err = json.Unmarshal(bytTxt, &data)

	if err != nil {
		fmt.Println("🔴", err)
		return err.Error()
	}
	return data["code"]
}
