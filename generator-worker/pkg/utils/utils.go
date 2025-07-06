package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"net/http"

	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/bikaxh/vid-gen/generator/pkg/prompts"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/joho/godotenv"

	"os/exec"
	"path/filepath"
)

type SceneGeneration struct {
	ClassName   string `json:"className"`
	Code        string `json:"code"`
	Description string `json:"description"`
	SceneTitle  string `json:"title"`
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

	fmt.Println("LLM Res", result)
	fmt.Println("parsed Res", *&data)

	// save to db
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

func CompileFile(fileName string) (*string, error) {

	pyFilePath := filepath.Join("..", "..", "final", "code", fileName+".py")
	pyAbsolutePath, err := filepath.Abs(pyFilePath)

	if err != nil {
		fmt.Println("Error getting absolute path for Python file:", err)
		return nil, err
	}

	mediaDirPath := filepath.Join("..", "..", "final", "media")
	mediaAbsolutePath, err := filepath.Abs(mediaDirPath)
	if err != nil {
		fmt.Println("Error getting absolute path for media dir:", err)
		return nil, err
	}

	cmd := exec.Command("manim", "-qh", pyAbsolutePath, fileName, "--media_dir", mediaAbsolutePath)
	_, err = cmd.CombinedOutput()

	if err != nil {

		return nil, err

	}

	videoPath := filepath.Join(mediaAbsolutePath, "videos", fileName, "1080p60", fileName+".mp4")
	return &videoPath, nil
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

		if success != nil {
			// update the code in db
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

func SaveSceneToDb(scene SceneGeneration, projectId string) {
	url := fmt.Sprintf("%s/create-scene/%s", GoDotEnvVariable("PRIMARY_SERVER"), projectId)
	fmt.Println("🟢", url)
	json, _ := json.Marshal(scene)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(json))

	if err != nil {

		fmt.Println("Got error while saving " + scene.SceneTitle)
		return
	}

	defer res.Body.Close()

	//  body, _ := io.ReadAll(res.Body)

	//  fmt.Println("REs body: ", string(body))

}

func ConcatUpload(projectId string) error {

	mediaDirPath := filepath.Join("..", "..", "final")
	mediaAbsolutePath, _ := filepath.Abs(mediaDirPath)

	// concat all
	//ffmpeg -f concat -safe 0 -i ffmpeg.txt -c copy ./final.mp4
	concatFilePath := filepath.Join(mediaAbsolutePath, "ffmpeg.txt")
	finalOutputPath := filepath.Join(mediaDirPath, projectId+".mp4")

	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", concatFilePath, "-c", "copy", finalOutputPath)
	_, err := cmd.CombinedOutput()

	if err != nil {

		fmt.Println(err.Error())
		return err
	}
	// upload

	cld, _ := cloudinary.NewFromURL(GoDotEnvVariable("CLOUDINARY_URL"))

	cld.Config.URL.Secure = true
	ctx := context.Background()

	res, err := cld.Upload.Upload(ctx, finalOutputPath, uploader.UploadParams{
		Folder:       "manimate",
		ResourceType: "video",
	})

	if err != nil {
		fmt.Println("🔴🔴", err)

	}

	// save video url
	reqBody := make(map[string]string)
	reqBody["videoUrl"] = res.SecureURL

	url := fmt.Sprintf("%s/save-video/%s", GoDotEnvVariable("PRIMARY_SERVER"), projectId)

	json, _ := json.Marshal(reqBody)

	response, _ := http.Post(url, "application/json", bytes.NewBuffer(json))

	defer response.Body.Close()
	return nil

}

var RedisClient *redis.Client

func ConnectRedisClient() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     GoDotEnvVariable("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("🟢 Connected to Redis")

}

func CleanResources() {
	fmt.Println("Clenaing")
	rootDirPath := filepath.Join("..", "..", "final")
	rootDirAbsolutePath, _ := filepath.Abs(rootDirPath)
	err := os.RemoveAll(rootDirAbsolutePath)
	if err != nil {
		fmt.Println("Error removing directory:", err)
	} else {
		fmt.Println("Directory removed successfully")
	}

	err = os.MkdirAll(rootDirAbsolutePath, 0755)
	err = os.MkdirAll(path.Join(rootDirAbsolutePath, "code"), 0755)
	err = os.MkdirAll(path.Join(rootDirAbsolutePath, "final"), 0755)
}

func WriteFFMPEGFile(className string) {
	FFMPEGFilePath := filepath.Join("..", "..", "final", "ffmpeg"+".txt")
	absolutePath, _ := filepath.Abs(FFMPEGFilePath)
	file, _ := os.OpenFile(absolutePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	mediaDirPath := filepath.Join("..", "..", "final", "media")
	mediaAbsolutePath, _ := filepath.Abs(mediaDirPath)
	videoPath := filepath.Join(mediaAbsolutePath, "videos", className, "1080p60", className+".mp4")

	pathToBeWritten := fmt.Sprintf("file '%s'\n", *&videoPath)

	file.WriteString(pathToBeWritten)
	file.Close()
}
