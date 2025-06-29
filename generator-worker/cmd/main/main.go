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
    "animationTypes": [
      "Write",
      "FadeIn",
      "Transform"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#81C784",
      "text": "#FFFFFF"
    },
    "instruction": "Introduce the concept of databases and then the idea of in-memory databases. Briefly mention Redis as a popular example.",
    "sceneTitle": "What is a Database?",
    "script": "Start with the title 'What is a Database?' written clearly. Show a stylized icon representing a traditional database (e.g., a cylinder). Then, introduce 'In-Memory Database' with text. Transition to a visual representing computer memory (RAM blocks). Finally, display the text 'Redis' prominently, perhaps with its logo, signifying it as a key player in this space.",
    "sequence": 1,
    "visualElements": [
      "Title Text: What is a Database?",
      "Icon: Traditional Database",
      "Text: In-Memory Database",
      "Visual: RAM Blocks",
      "Text/Logo: Redis"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "Transform",
      "Create"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#FFD54F",
      "text": "#E0E0E0"
    },
    "instruction": "Explain the concept of caching using a simple analogy. Focus on the benefit of speed and reducing repetitive work.",
    "sceneTitle": "Understanding Caching",
    "script": "Display the title 'The Power of Caching'. Use an analogy: show a person needing a tool. First, they walk to a large, distant storage shed (representing the main database). Show this path as slow. Then, show a small, easily accessible toolbox right next to the person holding the most frequently used tools (representing the cache). Illustrate the person quickly grabbing a tool from the toolbox. Emphasize the speed difference with text like 'Faster Access!'.",
    "sequence": 2,
    "visualElements": [
      "Title Text: The Power of Caching",
      "Visual: Person walking to distant storage",
      "Visual: Toolbox near the person",
      "Icon: Frequently Used Tool",
      "Text: Slow vs. Fast",
      "Animation: Arrow showing quick access"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "Transform",
      "MovingArrow"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#64B5F6",
      "text": "#FFFFFF"
    },
    "instruction": "Show how Redis acts as a cache layer between a web application and a database, speeding up data retrieval.",
    "sceneTitle": "Redis: Your Super-Fast Cache",
    "script": "Show a title 'Redis: Your Super-Fast Cache'. Display a flow: [User] -> [Web Server] -> [Main Database]. Animate data flowing slowly from the database. Then, introduce Redis into the flow: [User] -> [Web Server] -> [Redis Cache]. Show data being retrieved quickly from Redis. Add a condition: if data is NOT in Redis cache ([Redis Cache] -> [Main Database]), then show the data going back to Redis. Highlight the speed improvement.",
    "sequence": 3,
    "visualElements": [
      "Title Text: Redis: Your Super-Fast Cache",
      "Diagram Element: User",
      "Diagram Element: Web Server",
      "Diagram Element: Redis Cache (Box with 'Redis')",
      "Diagram Element: Main Database",
      "Arrows: Showing data flow direction"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "Create"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#BA68C8",
      "text": "#E0E0E0"
    },
    "instruction": "Briefly explain Redis's core data model: key-value pairs. Show simple examples of data stored.",
    "sceneTitle": "How Redis Stores Data",
    "script": "Display the title 'How Redis Stores Data'. Explain 'Key-Value Store'. Show examples appearing on screen: Text 'user:1' (key) with Text 'Alice' (value) next to it. Another example: Text 'session:xyz' (key) with a simple JSON-like object (value). Briefly mention other structures like lists or sets with simple visual representations if time permits.",
    "sequence": 4,
    "visualElements": [
      "Title Text: How Redis Stores Data",
      "Text: Key-Value Store",
      "Visual: Key 'user:1' -> Value 'Alice'",
      "Visual: Key 'session:xyz' -> Value { ... }",
      "Visual: Simple List representation"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "Transform",
      "DrawBorder"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#F06292",
      "text": "#FFFFFF"
    },
    "instruction": "Illustrate a real-world example of using Redis for caching a social media feed.",
    "sceneTitle": "Example: Social Media Feed",
    "script": "Show the title 'Example: Social Media Feed'. Display a simplified mobile feed UI mockup (scrolling posts). Show the app fetching feed data. Then, show Redis storing a pre-rendered version of the feed. When the user refreshes, the data is served quickly from Redis. Use arrows to show the data flow, emphasizing the path through Redis for speed. Add text like 'Reduced Load on Database'.",
    "sequence": 5,
    "visualElements": [
      "Title Text: Example: Social Media Feed",
      "Visual: Mockup of a social media feed",
      "Visual: Redis storing feed data (key: 'feed:user1')",
      "Animation: Arrow showing fast retrieval from Redis",
      "Text: Faster Feed Loading"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "Transform"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#FF8A65",
      "text": "#E0E0E0"
    },
    "instruction": "Provide another real-world example, such as caching product details or user sessions in e-commerce.",
    "sceneTitle": "Example: E-commerce",
    "script": "Display the title 'Example: E-commerce'. Show a simplified product page. Demonstrate fetching product details from a main database. Then, show Redis caching these details (e.g., key: 'product:123'). When another user views the same product, the details are served rapidly from Redis. Briefly show a user session being stored and retrieved from Redis, highlighting improved login/session management speed.",
    "sequence": 6,
    "visualElements": [
      "Title Text: Example: E-commerce",
      "Visual: Mockup of a product detail page",
      "Visual: Redis caching product details",
      "Visual: User session data in Redis",
      "Text: Quick Product Views",
      "Text: Stable User Sessions"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn",
      "GrowFromCenter"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#4FC3F7",
      "text": "#FFFFFF"
    },
    "instruction": "Summarize the key benefits of using Redis for caching: speed, reduced database load, and improved scalability.",
    "sceneTitle": "Why Redis for Caching?",
    "script": "Show the title 'Why Redis for Caching?'. Present three key benefits as distinct points, perhaps with icons: 1. Blazing Speed (Icon: Speedometer). 2. Reduced Database Load (Icon: Database with down arrow). 3. Scalability (Icon: Growing graph). Briefly elaborate on each point with concise text.",
    "sequence": 7,
    "visualElements": [
      "Title Text: Why Redis for Caching?",
      "Point 1: Blazing Speed (with Speedometer Icon)",
      "Point 2: Reduced Database Load (with DB Icon)",
      "Point 3: Scalability (with Graph Icon)"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#121212",
      "highlights": "#81C784",
      "text": "#FFFFFF"
    },
    "instruction": "Conclude the video by summarizing the main points: Redis is a fast in-memory database often used for caching to improve application performance.",
    "sceneTitle": "Summary & Key Takeaways",
    "script": "Display the title 'Summary'. Briefly recap: Redis is an in-memory data store. It excels at caching. Caching makes applications faster and reduces database strain. Show the Redis logo one last time. End with a concluding text like 'Boost Your Application Performance!'.",
    "sequence": 8,
    "visualElements": [
      "Title Text: Summary",
      "Text: Redis = Fast In-Memory Store",
      "Text: Caching improves performance",
      "Text: Reduces DB load",
      "Logo: Redis",
      "Final Text: Boost Your Application Performance!"
    ]
  }
]`)

	json.Unmarshal(jsonData, &scenes)

	var previousCode string

	for _, scene := range scenes {
		res, _ := utils.GenerateCode(scene, scenes, previousCode)
		fmt.Println(res.SceneTitle)
		// save to file
		utils.WriteToFile(res.ClassName, res.Code)
		// compile
		_, err := utils.CompileFile(res.ClassName)
		if err != nil {
			fmt.Println("Fixing 🔴", res.ClassName)
			utils.FixCode(fmt.Sprintf("%+v", err), *res)
		}

		pCode, _ := utils.ReadFile(res.ClassName)
		previousCode = *pCode

	}

}
