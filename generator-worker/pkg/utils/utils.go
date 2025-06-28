package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	// "strings"

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

func CompileFile(fileName string) {
	mediaPath := filepath.Join("..", "..", "final", "code", fileName+".py")
	absolutePath, err := filepath.Abs(mediaPath)

	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	fmt.Println(absolutePath)
	cmd := exec.Command("manim", "-qh", absolutePath, "/code", fileName+".py",fileName, "-- media_dir", absolutePath)

	res, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}
