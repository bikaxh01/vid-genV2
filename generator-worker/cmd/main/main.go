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
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Display the main title 'Understanding the Transformer Architecture'. Briefly introduce the Transformer as a powerful neural network model, highlighting its impact on Natural Language Processing (NLP) and other fields. Use a clean, modern font.",
    "sceneTitle": "Introduction to Transformers",
    "sequence": 1,
    "visualElements": [
      "Title Text: 'Understanding the Transformer Architecture'",
      "Subtitle Text: 'A Beginner's Guide'",
      "Simple Network Graphic (Optional)"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Transform",
      "FadeOut"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Briefly illustrate the limitations of older sequence models like RNNs. Show a sequence of words processed one by one, visually representing the 'sequential bottleneck' and mentioning challenges with long-range dependencies. Fade out this explanation.",
    "sceneTitle": "Limitations of RNNs/LSTMs",
    "sequence": 2,
    "visualElements": [
      "Text: 'Previous Models (RNNs/LSTMs)'",
      "Visual: Sequential Arrows (Word1 -> Word2 -> Word3)",
      "Text: 'Challenges: Sequential Processing, Long-Range Dependencies'"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Create",
      "Indicate"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#FFD700",
      "text": "#FFFFFF"
    },
    "instruction": "Introduce the core concept of 'Attention'. Use an analogy: when translating a sentence, we focus on specific words. Show a sentence where a target word is highlighted, and related source words are visually emphasized or connected.",
    "sceneTitle": "The Concept of Attention",
    "sequence": 3,
    "visualElements": [
      "Text: 'The Key Idea: Attention'",
      "Example Sentence Displayed",
      "Highlighting Effect on relevant words",
      "Connecting Arrows"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Transform",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Explain the Self-Attention mechanism. Show an input word embedding splitting into three distinct vectors: Query (Q), Key (K), and Value (V). Label each vector clearly.",
    "sceneTitle": "Self-Attention: Q, K, V Vectors",
    "sequence": 4,
    "visualElements": [
      "Text: 'Self-Attention'",
      "Input Word Representation (Box/Circle)",
      "Arrows splitting into Q, K, V vectors",
      "Labels: 'Query (Q)', 'Key (K)', 'Value (V)'"
    ]
  },
  {
    "animationTypes": [
      "Create",
      "Transform"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Demonstrate how attention scores are calculated. Show the Query vector of one word interacting (dot product) with the Key vectors of all words in the sequence. Visualize this interaction, perhaps with lines connecting Q to each K, possibly showing a similarity score calculation.",
    "sceneTitle": "Calculating Attention Scores",
    "sequence": 5,
    "visualElements": [
      "Q Vector",
      "K Vectors (for all words)",
      "Visualisation of Dot Product (e.g., lines with weights)",
      "Text: 'Score = Q * K^T'"
    ]
  },
  {
    "animationTypes": [
      "Transform",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#FFD700",
      "text": "#FFFFFF"
    },
    "instruction": "Explain the normalization step. Show the calculated scores being scaled (mention scaling factor briefly, e.g., divided by sqrt(d_k)) and then passed through a Softmax function. Visualize the scores transforming into probability weights (summing to 1).",
    "sceneTitle": "Normalization with Softmax",
    "sequence": 6,
    "visualElements": [
      "Attention Scores",
      "Scaling Operation Visualisation",
      "Softmax Function Visualisation",
      "Output: Attention Weights (probabilities)"
    ]
  },
  {
    "animationTypes": [
      "Transform",
      "Create"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Show the final step of self-attention. Multiply the attention weights by their corresponding Value vectors. Sum these weighted Value vectors to produce the final output vector for the initial word, representing its context-aware embedding.",
    "sceneTitle": "Weighted Sum of Values",
    "sequence": 7,
    "visualElements": [
      "Attention Weights",
      "Value Vectors (V)",
      "Multiplication Operation (Weights * V)",
      "Summation Operation",
      "Output: Contextualized Word Embedding"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Transform"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Introduce Multi-Head Attention. Explain that the Q, K, V process happens multiple times in parallel, each 'head' focusing on different aspects of the sequence relationships. Show the parallel paths and the final concatenation/projection.",
    "sceneTitle": "Multi-Head Attention",
    "sequence": 8,
    "visualElements": [
      "Text: 'Multi-Head Attention'",
      "Visual: Multiple parallel Q, K, V calculations",
      "Concatenation of results",
      "Final Projection Layer"
    ]
  },
  {
    "animationTypes": [
      "Create",
      "Write"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Describe the Encoder block. Show its components: Multi-Head Attention layer, followed by Add & Norm, then a Feed-Forward Network, and another Add & Norm layer. Indicate that multiple such blocks are typically stacked.",
    "sceneTitle": "The Encoder Block",
    "sequence": 9,
    "visualElements": [
      "Diagram: 'Encoder Block'",
      "Component 1: Multi-Head Attention",
      "Component 2: Add & Norm",
      "Component 3: Feed-Forward Network",
      "Component 4: Add & Norm",
      "Text: 'Stackable Layers'"
    ]
  },
  {
    "animationTypes": [
      "Create",
      "Write"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#FFD700",
      "text": "#FFFFFF"
    },
    "instruction": "Describe the Decoder block. Show its components: Masked Multi-Head Attention (preventing looking ahead), Add & Norm, Encoder-Decoder Attention (attending to encoder output), Add & Norm, Feed-Forward, Add & Norm. Indicate stacking.",
    "sceneTitle": "The Decoder Block",
    "sequence": 10,
    "visualElements": [
      "Diagram: 'Decoder Block'",
      "Component 1: Masked Multi-Head Attention",
      "Component 2: Add & Norm",
      "Component 3: Encoder-Decoder Attention",
      "Component 4: Add & Norm",
      "Component 5: Feed-Forward Network",
      "Component 6: Add & Norm",
      "Text: 'Stackable Layers'"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Transform"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Address the lack of sequential information in pure self-attention. Explain Positional Encoding: adding vectors that represent the position of each word in the sequence to the input embeddings. Show a visual representation.",
    "sceneTitle": "Positional Encoding",
    "sequence": 11,
    "visualElements": [
      "Text: 'Handling Word Order: Positional Encoding'",
      "Input Embedding + Positional Encoding Vector = Final Input",
      "Visual: Sine/Cosine waves or simple positional vectors"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "Create"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Show the complete Transformer architecture. Connect the input processing (embeddings + positional encoding), the stack of Encoder blocks, the stack of Decoder blocks, and the final output layer (e.g., a linear layer followed by Softmax for prediction).",
    "sceneTitle": "The Full Transformer Model",
    "sequence": 12,
    "visualElements": [
      "Diagram: Encoder Stack",
      "Diagram: Decoder Stack",
      "Input Layer (Embeddings + Positional Encoding)",
      "Output Layer (Linear + Softmax)",
      "Arrows showing data flow"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#FFD700",
      "text": "#FFFFFF"
    },
    "instruction": "List key applications where Transformers excel. Use icons or short text labels for examples like Machine Translation, Text Summarization, Question Answering, Chatbots, and even Computer Vision tasks.",
    "sceneTitle": "Applications of Transformers",
    "sequence": 13,
    "visualElements": [
      "Text: 'Applications'",
      "Icon/Label: Machine Translation",
      "Icon/Label: Text Summarization",
      "Icon/Label: Chatbots / Language Models",
      "Icon/Label: Computer Vision (Optional)"
    ]
  },
  {
    "animationTypes": [
      "Write",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#1A1A2E",
      "highlights": "#4D7FFF",
      "text": "#FFFFFF"
    },
    "instruction": "Conclude by summarizing the key takeaway: Self-Attention allows models to weigh the importance of different input parts, overcoming RNN limitations. Reiterate the significance and widespread use of Transformers.",
    "sceneTitle": "Conclusion and Impact",
    "sequence": 14,
    "visualElements": [
      "Text: 'Summary: Self-Attention is Key!'",
      "Text: 'Transformers revolutionized NLP and beyond.'",
      "Final graphic of the Transformer Architecture (simplified)"
    ]
  }
]`)

	json.Unmarshal(jsonData, &scenes)

	for _, scene := range scenes {

		res, _ := utils.GenerateCode(scene, scenes)
		fmt.Println(res.SceneTitle)
		// save to file
		utils.WriteToFile(res.ClassName, res.Code)
		// compile
		_, err := utils.CompileFile(res.ClassName)
		if err != nil {
			fmt.Println("Fixing 🔴", res.ClassName)
			utils.FixCode(fmt.Sprintf("%+v", err), *res)
		}
	}

}
