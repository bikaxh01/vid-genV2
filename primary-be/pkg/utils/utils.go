package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/genai"
)

func HashPassword(password string) (*string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return nil, err
	}

	hash := string(bytes)

	return &hash, nil
}

func ComparePassword(hashPW, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashPW), []byte(password))

	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateToken(email, id string) (*string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := GoDotEnvVariable("JWT_SECRET")

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &token, nil
}

func VerifyToken(tokenString string) (*string, error) {
	secretKey := GoDotEnvVariable("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Error while validating token")
	}

	if !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Error while validating token")
	}
	userId, _ := claims["user_id"].(string)
	return &userId, nil
}

func GoDotEnvVariable(key string) string {

	err := godotenv.Load("../../.env")

	if err != nil {

		log.Fatalf("Error loading .env file", err)
	}

	return os.Getenv(key)
}

// type SceneType struct {
// 	AnimationTypes []string          `json:"animationTypes"`
// 	ColorScheme    map[string]string `json:"colorScheme"`
// 	Instruction    string            `json:"instruction"`
// 	SceneTitle     string            `json:"sceneTitle"`
// 	VisualElements []string          `json:"visualElements"`
// }

// type GeneratePlanType []SceneType

func GeneratePlan(prompt string) (map[string]any, error) {

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
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"metaData": {
					Type:        genai.TypeObject,
					Description: "This will contain the title and description of the Topic",
					Properties: map[string]*genai.Schema{
						"title": {
							Type:        genai.TypeString,
							Description: "This will be the title of the topic",
						},
						"description": {
							Type:        genai.TypeString,
							Description: "This will be the short description of the topic",
						},
					},
					Required: []string{"title", "description"},
				},
				"scenesData": {
					Type:        genai.TypeArray,
					Description: "This will be the array of scenes",
					Items: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"sceneTitle": {
								Type:        genai.TypeString,
								Description: "The title of the scene, describing the main focus or idea.",
							},
							"script": {
								Type:        genai.TypeString,
								Description: "The in detail script  for  the scene.",
							},
							"instruction": {
								Type:        genai.TypeString,
								Description: "A detailed explanation of what happens in the scene.",
							},
							"visualElements": {
								Type:        genai.TypeArray,
								Description: "List of visual elements and components shown in the scene (e.g., diagrams, titles, arrows).",
								Items: &genai.Schema{
									Type: genai.TypeString,
								},
							},
							"sequence": {
								Type:        genai.TypeNumber,
								Description: "The sequence of the scene",
							},
							"colorScheme": {
								Type:        genai.TypeObject,
								Description: "A consistent color scheme used in the scene.",
								Properties: map[string]*genai.Schema{
									"background": {
										Type:        genai.TypeString,
										Description: "The background color (e.g., 'light blue')",
									},
									"text": {
										Type:        genai.TypeString,
										Description: "The primary text color (e.g., 'white')",
									},
									"highlights": {
										Type:        genai.TypeString,
										Description: "The color used for highlighting key elements (e.g., 'yellow')",
									},
								},
								Required: []string{"background", "text", "highlights"},
							},
							"animationTypes": {
								Type:        genai.TypeArray,
								Description: "List of animation types used in the scene (e.g., Write, FadeIn, Transform).",
								Items: &genai.Schema{
									Type: genai.TypeString,
								},
							},
						},
						Required: []string{
							"sceneTitle",
							"instruction",
							"visualElements",
							"colorScheme",
							"script",
							"sequence",
							"animationTypes",
						},
					},
				},
			},
			Required: []string{"metaData", "scenesData"},
		},
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"models/gemini-2.5-flash-lite-preview-06-17",
		genai.Text(GetPrompt(prompt)),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	res := []byte(result.Text())

	var data map[string]interface{}
	err = json.Unmarshal(res, &data)
      
	if err != nil {
		fmt.Println("🔴", err)
		return nil, err
	}

	return data, nil

}
