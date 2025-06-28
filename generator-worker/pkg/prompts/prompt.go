package prompts

import (
	"encoding/json"
	"fmt"
)

type Scene struct {
	AnimationTypes []string          `json:"animationTypes"`
	ColorScheme    map[string]string `json:"colorScheme"`
	Instruction    string            `json:"instruction"`
	SceneTitle     string            `json:"sceneTitle"`
	VisualElements []string          `json:"visualElements"`
	Sequence       int               `json:"sequence"`
}

type Scenes []Scene

func GetSceneGenerationPrompt(sceneMetadata Scene, scenes Scenes) string {

	colorBytes, _ := json.Marshal(sceneMetadata.ColorScheme)
	animationBytes, _ := json.Marshal(sceneMetadata.AnimationTypes)
	visualBytes, _ := json.Marshal(sceneMetadata.VisualElements)

	colorScheme := string(colorBytes)
	animationTypes := string(animationBytes)
	visualElements := string(visualBytes)

	jsonData, _ := json.Marshal(scenes)

	var prompt string = fmt.Sprintf(`You are an expert-level AI assistant specializing in **Python animation development** using **Manim Community Edition (v0.19.0)**. Your task is to generate Python code for a **single animation scene** based strictly on a provided configuration.

---

**Your Objective**:
- Generate a **single, production-ready Manim scene class** that visually expresses the configuration described in the section titled **"Scene to Generate"**.
- **Do not** include content or logic from other scenes listed in the broader context.
- All visual elements, animations,Instruction and structure must adhere strictly to the configuration provided.

---

 **Rules and Constraints**:
- Assume a 16:9 screen layout (1920x1080 resolution).
- Use **only this import** at the top:
  \from manim import *\

- Define **only one** scene class.

- The class **must follow this naming format**:
  \<sequence>_<PascalCaseTitle>\  
  Example: \S01_IntroductionToGravity\

- The class must use only **Manim CE v0.19.0-compatible syntax**.

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

 **Scene to Generate**  
Use this section ONLY to build your Manim scene:

- **Title**: %v  
- **Sequence**: %v  
- **Instruction**: %v  
- **Color Scheme**: %v  
- **Animation Types**: %v  
- **Visual Elements**: %v

---

 **Final Checklist Before Output**:
-  Only one class  
-  Only from manim import * at the top  
-  Strict adherence to the given scene configuration  
-  No template code, boilerplate, or external logic  
-  Bug-free and executable in **Manim CE v0.19.0**

Begin now. Output only the Manim scene class.`, string(jsonData), sceneMetadata.SceneTitle, sceneMetadata.Sequence, sceneMetadata.Instruction, colorScheme, animationTypes, visualElements)

	return prompt
}

// **Code Class Requirements** (for Code(...)):
// When using the \Code\ object, use only these parameters and their correct data types:

// \\\python
// Code(
//   code_file: str | None = None,
//   code_string: str | None = None,
//   language: str | None = None,
//   formatter_style: str = "vim",
//   tab_width: int = 4,
//   add_line_numbers: bool = True,
//   line_numbers_from: int = 1,
//   background: Literal["rectangle", "window"] = "rectangle",
//   background_config: dict | None = None
// )
// \\\


func GetFixCodePrompt () string{

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
-  Return **only** the corrected class (with \from manim import *\ at the top) — nothing else`)
return prompt
}