package prompts

import (
	"encoding/json"
	"fmt"
)

type Scene struct {
	AnimationTypes []string          `json:"animationTypes"`
	ColorScheme    map[string]string `json:"colorScheme"`
	Description    string            `json:"description"`
	Script         string            `json:"script"`
	SceneTitle     string            `json:"sceneTitle"`
	VisualElements []string          `json:"visualElements"`
	Sequence       int               `json:"sequence"`
}

type Scenes []Scene

func GetSceneGenerationSystemPrompt(scenes Scenes) string {

	jsonData, _ := json.Marshal(scenes)

	var prompt string = fmt.Sprintf(`You are an expert-level AI assistant specializing in **Python animation development** using **Manim Community Edition**. Your task is to generate Python code for a **single animation scene** based strictly on a provided configuration.

---

**Your Objective**:
- Generate a **single, production-ready Manim scene class** that visually expresses the configuration described in the section titled **"Scene to Generate"**.
- **Do not** include content or logic from other scenes listed in the broader context.
- All visual elements, animations,Instruction and structure must adhere strictly to the configuration provided.
- Make sure the text should not overlap 
- text and model should be at there correct position
---

 **Rules and Constraints**:
- Assume a 16:9 screen layout (1920x1080 resolution).
- Make sure the visual elements you are adding make sense for the scene
- all the text and animation should be inside the frame
- Use **only this import** at the top:
  \from manim import *\

- Define **only one** scene class.

- The class **must follow this naming format**:
  \<sequence>_<PascalCaseTitle>\  
  Example: \S01_IntroductionToGravity\

- The class must use only **Manim CE compatible syntax**.

- The code should be:
  - Clean and efficient  
  - Free of deprecated methods  
  - Self-contained and runnable

- **Allowed animation types**: Use only those explicitly mentioned under \Animation Types\.

- **Allowed color codes**: Use only the hex values provided in the color scheme (e.g., \#FFFFFF\). Do **not** use named colors.

- **Visual Elements**: Use only those described under \Visual Elements\ — interpret and render them as accurately as possible using Manim objects (e.g., Text, MathTex, Rectangle, Arrow, etc.).

- **Scene Layout**:
  - Maintain proper spacing and alignment  
  - Ensure all text, shapes, and transitions are **clearly visible** and **centered or well-positioned**  
  - Avoid visual clutter

- **Prohibited**:
  - Third-party libraries  
  - External assets (SVGs, images, fonts)  
  - References to web links, file paths, or contacts  
  - Deprecated utilities 
---


 **Scene Configuration Context**  
(*Do NOT use this section directly. It is provided only for background understanding.*)
\\\
%v
\\\
---
 **Final Checklist Before Output**:
-  Only one class  
-  Only from manim import * at the top  
-  Strict adherence to the given scene configuration  
-  No template code, boilerplate, or external logic  
-  Bug-free and executable in **Manim CE v0.19.0**

---

OUTPUT_FORMAT
{
"sceneTitle": Title of the scene (string),
"description": short description of the scene (string),
"className": className of the scene (string),
"code": Code of the scene (string),
}


Note: output in JSON format with keys: "sceneTitle” (string), "description” (string), and "className” (string) "code" (string).
NOTE : Output should not be in markdown format


Output_Example
{
"sceneTitle": "Introduction to Redis",
"description": "Introduction to the topic. The title 'Redis: The In-Memory Data Store' appears on screen, followed by a brief subtitle explaining its nature as a fast, open-source, in-memory data structure store.",
"className": "S01_IntroductionToRedis",
"code": "from manim import *\n\nclass S01_IntroductionToRedis(Scene):\n    def construct(self):\n        # Set background color\n        self.camera.background_color = \"#FFFFFF\"\n        \n        # Create title text\n        title = Text(\n            \"Redis: The In-Memory Data Store\",\n            font_size=48,\n            color=\"#333333\",\n            weight=BOLD\n        ).shift(UP * 1.5)\n        \n        # Create subtitle text\n        subtitle = Text(\n            \"A Fast, Open-Source, In-Memory Data Structure Store\",\n            font_size=32,\n            color=\"#333333\"\n        ).shift(DOWN * 0.5)\n        \n        # Create Redis logo (stylized red bird)\n        redis_logo = VGroup()\n        \n        # Bird body (circle)\n        body = Circle(\n            radius=0.8,\n            color=\"#FF6347\",\n            fill_opacity=1\n        )\n        \n        # Bird beak (triangle)\n        beak = Triangle(\n            color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.3).rotate(-PI/2).shift(LEFT * 0.8)\n        \n        # Bird eye\n        eye = Dot(\n            point=UP * 0.2 + LEFT * 0.2,\n            radius=0.1,\n            color=\"#FFFFFF\"\n        )\n        \n        # Bird tail (small triangles)\n        tail1 = Triangle(\n            color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.4).rotate(PI/4).shift(RIGHT * 0.6 + DOWN * 0.3)\n        \n        tail2 = Triangle(\n          
  color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.4).rotate(PI/6).shift(RIGHT * 0.7 + UP * 0.1)\n        \n        # Assemble the logo\n        redis_logo.add(body, beak, eye, tail1, tail2)\n        redis_logo.scale(0.8).shift(LEFT * 5 + UP * 1.5)\n        \n        # Animate the scene\n        # Title fades in\n        self.play(FadeIn(title), run_time=2)\n        \n        # Subtitle appears with Write animation\n        self.play(Write(subtitle), run_time=2)\n        \n        # Redis logo fades in next to the title\n        self.play(FadeIn(redis_logo), run_time=1.5)\n        \n        # Hold the final frame\n        self.wait(2)"
}

`, string(jsonData))

	return prompt
}

func GetSceneGenerationUserPrompt(previousCode string, sceneMetadata Scene) string {
	colorBytes, _ := json.Marshal(sceneMetadata.ColorScheme)
	animationBytes, _ := json.Marshal(sceneMetadata.AnimationTypes)
	visualBytes, _ := json.Marshal(sceneMetadata.VisualElements)

	colorScheme := string(colorBytes)
	animationTypes := string(animationBytes)
	visualElements := string(visualBytes)

	prompt := fmt.Sprintf(`
  
  ---
**Previous Scene Code***

%v
---
 **Scene to Generate**
Use this section ONLY to build your Manim scene:

- **Title**: %v
- **Sequence**: %v
- **Description**: %v
- **script**: %v
- **Color Scheme**: %v
- **Animation Types**: %v
- **Visual Elements**: %v
  
  
  `, previousCode, sceneMetadata.SceneTitle, sceneMetadata.Sequence, sceneMetadata.Description, sceneMetadata.Script, colorScheme, animationTypes, visualElements)

	return prompt
}

func GetFixCodePrompt() string {

	prompt := fmt.Sprintf(`
		You are an expert AI Python developer with advanced knowledge of **Manim Community Edition (v0.19.0)**. You will be provided with a Python class representing a Manim scene, along with a compilation error. Your task is to **analyze the error and fix only what is necessary**.

---

 **Your Objective**:
- Correct the **exact compilation error** provided.
- Ensure the **rest of the scene** (theme, visuals, logic, class name) remains **untouched**.
- Your fix should be clean, minimal, and **production-ready**.

---

**Strict Rules**:

- Use **only this import** at the top:
  \from manim import *\

- Do **not**:
  - Change the class name
  - Add or remove visual/animation logic unless essential for the fix
  - Use deprecated or ManimGL methods
  - Use any third-party libraries
  - Include external assets (SVGs, images, etc.)
  - Use color **names** (use **hex codes only**)

- Your fix **must** be:
  - Fully compatible with **Manim CE v0.19.0**
  - Clean, readable, and correct
  - A **single self-contained scene class**
---

**Final Checklist**:

-  Only one class  
-  Preserve original logic and layout  
-  Fix **only** what causes the error  
-  Compatible with **Manim CE v0.19.0**  
-  Return **only** the corrected class (with \from manim import *\ at the top) — nothing else


OUTPUT_FORMAT
{
"code": fixed Code of the scene (string),
}

Note: output in JSON format with keys:  "code" (string).
NOTE : Output should not be in markdown format

Output_Example
{
"code": "from manim import *\n\nclass S01_IntroductionToRedis(Scene):\n    def construct(self):\n        # Set background color\n        self.camera.background_color = \"#FFFFFF\"\n        \n        # Create title text\n        title = Text(\n            \"Redis: The In-Memory Data Store\",\n            font_size=48,\n            color=\"#333333\",\n            weight=BOLD\n        ).shift(UP * 1.5)\n        \n        # Create subtitle text\n        subtitle = Text(\n            \"A Fast, Open-Source, In-Memory Data Structure Store\",\n            font_size=32,\n            color=\"#333333\"\n        ).shift(DOWN * 0.5)\n        \n        # Create Redis logo (stylized red bird)\n        redis_logo = VGroup()\n        \n        # Bird body (circle)\n        body = Circle(\n            radius=0.8,\n            color=\"#FF6347\",\n            fill_opacity=1\n        )\n        \n        # Bird beak (triangle)\n        beak = Triangle(\n            color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.3).rotate(-PI/2).shift(LEFT * 0.8)\n        \n        # Bird eye\n        eye = Dot(\n            point=UP * 0.2 + LEFT * 0.2,\n            radius=0.1,\n            color=\"#FFFFFF\"\n        )\n        \n        # Bird tail (small triangles)\n        tail1 = Triangle(\n            color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.4).rotate(PI/4).shift(RIGHT * 0.6 + DOWN * 0.3)\n        \n        tail2 = Triangle(\n          
  color=\"#FF6347\",\n            fill_opacity=1\n        ).scale(0.4).rotate(PI/6).shift(RIGHT * 0.7 + UP * 0.1)\n        \n        # Assemble the logo\n        redis_logo.add(body, beak, eye, tail1, tail2)\n        redis_logo.scale(0.8).shift(LEFT * 5 + UP * 1.5)\n        \n        # Animate the scene\n        # Title fades in\n        self.play(FadeIn(title), run_time=2)\n        \n        # Subtitle appears with Write animation\n        self.play(Write(subtitle), run_time=2)\n        \n        # Redis logo fades in next to the title\n        self.play(FadeIn(redis_logo), run_time=1.5)\n        \n        # Hold the final frame\n        self.wait(2)"
}

`)
	return prompt
}
