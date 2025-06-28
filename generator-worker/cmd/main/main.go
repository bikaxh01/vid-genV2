package main

import (
	"encoding/json"
	"fmt"

	"github.com/bikaxh/vid-gen/generator/pkg/prompts"
	"github.com/bikaxh/vid-gen/generator/pkg/utils"
)

func main() {
	// get project id

	// get project plan

	// map over it

	var scenes prompts.Scenes

	jsonData := []byte(`[
    {
      "animationTypes": ["Write", "FadeIn"],
      "colorScheme": {
        "background": "#1E1E1E",
        "highlights": "#FFD700",
        "text": "#FFFFFF"
      },
      "instruction":
        "Display the title 'Apache Kafka: A Beginner's Guide' prominently in the center of the screen. Fade in the subtitle 'Understanding the Fundamentals'. The background is a dark, professional gray.",
      "sceneTitle": "Introduction to Kafka",
      "sequence": 1,
      "visualElements": [
        "Title: Apache Kafka: A Beginner's Guide",
        "Subtitle: Understanding the Fundamentals"
      ]
    },
    {
      "animationTypes": ["Write", "Create", "Arrow"],
      "colorScheme": {
        "background": "#1E1E1E",
        "highlights": "#00CED1",
        "text": "#FFFFFF"
      },
      "instruction":
        "Introduce the problem Kafka solves: traditional messaging systems being slow and complex. Show a simple diagram with a Producer sending a message to a Consumer, but illustrate it as a slow, direct, point-to-point connection with a bottleneck. Use arrows to show data flow.",
      "sceneTitle": "The Problem: Traditional Messaging",
      "sequence": 2,
      "visualElements": [
        "Text: Traditional Messaging",
        "Diagram: Producer -> Consumer (Point-to-Point)",
        "Arrow: Slow Data Transfer",
        "Text: Bottlenecks & Complexity"
      ]
    },
    {
      "animationTypes": ["Write", "Create", "Transform", "FadeIn"],
      "colorScheme": {
        "background": "#1E1E1E",
        "highlights": "#32CD32",
        "text": "#FFFFFF"
      },
      "instruction":
        "Introduce Kafka as the solution. Show a central Kafka cluster acting as a message broker. Producers send messages to Kafka, and Consumers read from Kafka. Use distinct icons for Producers, Kafka Cluster, and Consumers. Show data flowing from Producers to Kafka, and then from Kafka to Consumers.",
      "sceneTitle": "Introducing Kafka: The Solution",
      "sequence": 3,
      "visualElements": [
        "Text: Apache Kafka",
        "Icon: Producer",
        "Icon: Kafka Cluster",
        "Icon: Consumer",
        "Arrows: Producer -> Kafka, Kafka -> Consumer",
        "Text: High Throughput, Scalable, Fault-Tolerant"
      ]
    }
  ]`)

	json.Unmarshal(jsonData, &scenes)

	res, _ := utils.GenerateCode(scenes[1], scenes)
	fmt.Println(res.SceneTitle)
	// save to file
	utils.WriteToFile(res.ClassName, res.Code)
	// compile
	utils.CompileFile(res.ClassName)
}
