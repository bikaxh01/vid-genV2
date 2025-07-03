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
      "FadeIn",
      "Write"
    ],
    "colorScheme": {
      "background": "#FFFFFF",
      "highlights": "#FFD700",
      "text": "#000000"
    },
    "description": "Introduce the concept of Neural Networks and their inspiration from the human brain.",
    "sceneTitle": "What is a Neural Network?",
    "script": "Start with a title 'What is a Neural Network?'. Then, show a simplified graphic of the human brain. Fade in text that reads 'Inspired by the human brain'. Transition to a visual representation of interconnected nodes, similar to neurons. Explain that a neural network is a computing system inspired by biological neural networks.",
    "sequence": 1,
    "visualElements": [
      "Title: What is a Neural Network?",
      "Graphic: Simplified Human Brain",
      "Text: Inspired by the human brain",
      "Graphic: Interconnected Nodes (Neurons)"
    ]
  },
  {
    "animationTypes": [
      "Create",
      "Transform"
    ],
    "colorScheme": {
      "background": "#E0F7FA",
      "highlights": "#0077B6",
      "text": "#023E8A"
    },
    "description": "Break down the basic structure of an artificial neural network: input layer, hidden layer(s), and output layer.",
    "sceneTitle": "The Anatomy of a Neural Network",
    "script": "Display the title 'The Anatomy of a Neural Network'. Show three distinct vertical layers of circles. Label the leftmost layer 'Input Layer'. Label the middle layer 'Hidden Layer'. Label the rightmost layer 'Output Layer'. Draw lines connecting each node in the input layer to every node in the hidden layer, and similarly connecting the hidden layer to the output layer. Briefly explain the role of each layer: input takes data, hidden layers process it, and output gives the result.",
    "sequence": 2,
    "visualElements": [
      "Title: The Anatomy of a Neural Network",
      "Graphic: Input Layer (Circles)",
      "Graphic: Hidden Layer (Circles)",
      "Graphic: Output Layer (Circles)",
      "Lines: Connections between layers",
      "Text Labels: Input Layer, Hidden Layer, Output Layer"
    ]
  },
  {
    "animationTypes": [
      "Transform",
      "Write"
    ],
    "colorScheme": {
      "background": "#FFF9C4",
      "highlights": "#FBC02D",
      "text": "#424242"
    },
    "description": "Explain how information flows through the network and how weights and biases influence the output.",
    "sceneTitle": "How Information Flows",
    "script": "Focus on a single neuron in the hidden layer. Show input values (represented by numbers or small circles) flowing into it from the input layer via weighted connections. Illustrate a 'weight' associated with each connection, perhaps as a number that changes. Show these weighted inputs summing up. Add a 'bias' term. Display an 'activation function' icon (like a sigmoid curve). Show the output of the activation function becoming the neuron's output, which then feeds into the next layer. Explain that weights and biases are adjusted during training.",
    "sequence": 3,
    "visualElements": [
      "Graphic: Single Neuron",
      "Input Values",
      "Weighted Connections (with numbers)",
      "Bias Term (number)",
      "Activation Function Graphic",
      "Output Value",
      "Text: Weights",
      "Text: Bias",
      "Text: Activation Function"
    ]
  },
  {
    "animationTypes": [
      "Transform",
      "Indicate",
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#E8F5E9",
      "highlights": "#388E3C",
      "text": "#1B5E20"
    },
    "description": "Illustrate the concept of training a neural network using a simple example, like recognizing a handwritten digit.",
    "sceneTitle": "How Neural Networks Learn (Training)",
    "script": "Show the input layer receiving an image of a handwritten digit (e.g., '7'). The network processes it and produces an output (e.g., it guesses '1'). Display the correct answer ('7') and highlight the difference (the error). Introduce the concept of 'backpropagation' with an arrow going backward through the network. Show the weights and biases adjusting slightly to reduce the error. Repeat this process a few times, showing the network's guesses getting closer to the correct answer, indicating learning.",
    "sequence": 4,
    "visualElements": [
      "Graphic: Handwritten Digit Image (Input)",
      "Network Structure",
      "Output Guess (e.g., '1')",
      "Correct Answer (e.g., '7')",
      "Error Calculation",
      "Arrow: Backpropagation",
      "Animated Weights/Biases Adjustment",
      "Text: Training",
      "Text: Backpropagation",
      "Text: Error Reduction"
    ]
  },
  {
    "animationTypes": [
      "FadeIn",
      "Write",
      "Transform"
    ],
    "colorScheme": {
      "background": "#FCE4EC",
      "highlights": "#E91E63",
      "text": "#880E4F"
    },
    "description": "Provide real-world examples of where neural networks are used.",
    "sceneTitle": "Real-World Applications",
    "script": "Display the title 'Real-World Applications'. Show icons or simple graphics representing different applications: a smartphone for voice assistants (like Siri/Alexa), a car for self-driving vehicles, a shopping cart for recommendation systems (like Netflix/Amazon), and an image of a medical scan for image recognition in healthcare. Briefly describe each application and how neural networks enable it.",
    "sequence": 5,
    "visualElements": [
      "Title: Real-World Applications",
      "Icon: Smartphone (Voice Assistant)",
      "Icon: Self-Driving Car",
      "Icon: Shopping Cart (Recommendations)",
      "Graphic: Medical Scan (Image Recognition)",
      "Text Descriptions for each application"
    ]
  },
  {
    "animationTypes": [
      "FadeIn"
    ],
    "colorScheme": {
      "background": "#FFFFFF",
      "highlights": "#4CAF50",
      "text": "#000000"
    },
    "description": "Conclude the video by summarizing the key concepts and encouraging further learning.",
    "sceneTitle": "Conclusion & Next Steps",
    "script": "Show the title 'Conclusion'. Briefly recap the main points: what a neural network is, its basic structure (layers, neurons, weights, biases), how it learns (training, backpropagation), and its applications. Display a friendly message like 'Neural Networks are powerful tools!' Fade in text encouraging viewers to learn more.",
    "sequence": 6,
    "visualElements": [
      "Title: Conclusion",
      "Summary Text Points",
      "Encouragement Text"
    ]
  }
]`)

	json.Unmarshal(jsonData, &scenes)

	// var previousCode string

	var generatedScenesMetaData []utils.SceneGeneration = []utils.SceneGeneration{
		{
			ClassName: "S01_WhatIsANeuralNetwork",
			Code: `from manim import *

class S01_WhatIsANeuralNetwork(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#FFFFFF"
        
        # Create title
        title = Text(
            "What is a Neural Network?",
            font_size=56,
            color="#000000",
            weight=BOLD
        ).shift(UP * 2.5)
        
        # Create simplified brain graphic
        brain = VGroup()
        
        # Brain outline (two hemispheres)
        left_hemisphere = Ellipse(
            width=2.5,
            height=3,
            color="#FFD700",
            fill_opacity=0.3,
            stroke_width=4
        ).shift(LEFT * 0.6)
        
        right_hemisphere = Ellipse(
            width=2.5,
            height=3,
            color="#FFD700",
            fill_opacity=0.3,
            stroke_width=4
        ).shift(RIGHT * 0.6)
        
        # Brain stem
        brain_stem = Rectangle(
            width=1,
            height=1.5,
            color="#FFD700",
            fill_opacity=0.3,
            stroke_width=4
        ).shift(DOWN * 2)
        
        # Brain folds (simplified)
        fold1 = Arc(
            radius=0.8,
            start_angle=PI/4,
            angle=PI/2,
            color="#FFD700",
            stroke_width=3
        ).shift(LEFT * 0.5 + UP * 0.5)
        
        fold2 = Arc(
            radius=0.8,
            start_angle=-PI/4,
            angle=PI/2,
            color="#FFD700",
            stroke_width=3
        ).shift(RIGHT * 0.5 + UP * 0.3)
        
        fold3 = Arc(
            radius=0.6,
            start_angle=PI/3,
            angle=PI/3,
            color="#FFD700",
            stroke_width=3
        ).shift(DOWN * 0.5)
        
        brain.add(left_hemisphere, right_hemisphere, brain_stem, fold1, fold2, fold3)
        brain.scale(0.8).shift(LEFT * 3.5)
        
        # Inspiration text
        inspiration_text = Text(
            "Inspired by the human brain",
            font_size=36,
            color="#000000",
            slant=ITALIC
        ).shift(LEFT * 3.5 + DOWN * 3)
        
        # Neural network visualization
        neural_net = VGroup()
        
        # Create nodes
        layer1_nodes = VGroup(*[
            Circle(radius=0.25, color="#FFD700", fill_opacity=0.8, stroke_width=2)
            for _ in range(3)
        ]).arrange(DOWN, buff=0.8)
        
        layer2_nodes = VGroup(*[
            Circle(radius=0.25, color="#FFD700", fill_opacity=0.8, stroke_width=2)
            for _ in range(4)
        ]).arrange(DOWN, buff=0.6)
        
        layer3_nodes = VGroup(*[
            Circle(radius=0.25, color="#FFD700", fill_opacity=0.8, stroke_width=2)
            for _ in range(2)
        ]).arrange(DOWN, buff=0.8)
        
        # Position layers
        layer1_nodes.shift(LEFT * 2)
        layer2_nodes.shift(ORIGIN)
        layer3_nodes.shift(RIGHT * 2)
        
        # Create connections
        connections = VGroup()
        for n1 in layer1_nodes:
            for n2 in layer2_nodes:
                line = Line(
                    n1.get_center(),
                    n2.get_center(),
                    color="#000000",
                    stroke_width=1,
                    stroke_opacity=0.5
                )
                connections.add(line)
        
        for n2 in layer2_nodes:
            for n3 in layer3_nodes:
                line = Line(
                    n2.get_center(),
                    n3.get_center(),
                    color="#000000",
                    stroke_width=1,
                    stroke_opacity=0.5
                )
                connections.add(line)
        
        neural_net.add(connections, layer1_nodes, layer2_nodes, layer3_nodes)
        neural_net.scale(0.7).shift(RIGHT * 3 + DOWN * 0.5)
        
        # Explanation text
        explanation = Text(
            "A computing system inspired by\nbiological neural networks",
            font_size=28,
            color="#000000",
            line_spacing=1.2
        ).shift(DOWN * 3)
        
        # Animations
        self.play(Write(title), run_time=2)
        self.wait(1)
        
        self.play(FadeIn(brain), run_time=2)
        self.wait(0.5)
        
        self.play(FadeIn(inspiration_text), run_time=1.5)
        self.wait(1)
        
        self.play(
            FadeIn(neural_net),
            brain.animate.shift(UP * 0.5),
            inspiration_text.animate.shift(UP * 0.5),
            run_time=2
        )
        self.wait(0.5)
        
        self.play(Write(explanation), run_time=2)
        self.wait(2)`, Description: "Hello wok", SceneTitle: "Introduction"},
		{
			ClassName: "S02_TheAnatomyOfANeuralNetwork",
			Code: `from manim import *

class S02_TheAnatomyOfANeuralNetwork(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#E0F7FA"
        
        # Create title
        title = Text(
            "The Anatomy of a Neural Network",
            font_size=48,
            color="#023E8A",
            weight=BOLD
        ).shift(UP * 3.2)
        
        # Create layers of nodes
        # Input layer - 4 nodes
        input_nodes = VGroup(*[
            Circle(
                radius=0.3,
                color="#0077B6",
                fill_opacity=0.8,
                stroke_width=3,
                stroke_color="#023E8A"
            )
            for _ in range(4)
        ]).arrange(DOWN, buff=0.6)
        input_nodes.shift(LEFT * 4)
        
        # Hidden layer - 5 nodes
        hidden_nodes = VGroup(*[
            Circle(
                radius=0.3,
                color="#0077B6",
                fill_opacity=0.8,
                stroke_width=3,
                stroke_color="#023E8A"
            )
            for _ in range(5)
        ]).arrange(DOWN, buff=0.5)
        hidden_nodes.shift(ORIGIN)
        
        # Output layer - 3 nodes
        output_nodes = VGroup(*[
            Circle(
                radius=0.3,
                color="#0077B6",
                fill_opacity=0.8,
                stroke_width=3,
                stroke_color="#023E8A"
            )
            for _ in range(3)
        ]).arrange(DOWN, buff=0.6)
        output_nodes.shift(RIGHT * 4)
        
        # Create layer labels
        input_label = Text(
            "Input Layer",
            font_size=28,
            color="#023E8A",
            weight=BOLD
        ).next_to(input_nodes, DOWN, buff=0.5)
        
        hidden_label = Text(
            "Hidden Layer",
            font_size=28,
            color="#023E8A",
            weight=BOLD
        ).next_to(hidden_nodes, DOWN, buff=0.5)
        
        output_label = Text(
            "Output Layer",
            font_size=28,
            color="#023E8A",
            weight=BOLD
        ).next_to(output_nodes, DOWN, buff=0.5)
        
        # Create connections
        input_to_hidden_connections = VGroup()
        for input_node in input_nodes:
            for hidden_node in hidden_nodes:
                connection = Line(
                    input_node.get_center(),
                    hidden_node.get_center(),
                    color="#0077B6",
                    stroke_width=1.5,
                    stroke_opacity=0.6
                )
                input_to_hidden_connections.add(connection)
        
        hidden_to_output_connections = VGroup()
        for hidden_node in hidden_nodes:
            for output_node in output_nodes:
                connection = Line(
                    hidden_node.get_center(),
                    output_node.get_center(),
                    color="#0077B6",
                    stroke_width=1.5,
                    stroke_opacity=0.6
                )
                hidden_to_output_connections.add(connection)
        
        # Create explanation text
        explanation = VGroup(
            Text("Input: Takes data", font_size=22, color="#023E8A"),
            Text("Hidden: Processes information", font_size=22, color="#023E8A"),
            Text("Output: Gives the result", font_size=22, color="#023E8A")
        ).arrange(DOWN, buff=0.3, aligned_edge=LEFT)
        explanation.shift(DOWN * 3.2)
        
        # Animations
        self.play(Create(title), run_time=1.5)
        self.wait(0.5)
        
        # Create input layer
        self.play(
            Create(input_nodes),
            Create(input_label),
            run_time=1.5
        )
        self.wait(0.5)
        
        # Create hidden layer
        self.play(
            Create(hidden_nodes),
            Create(hidden_label),
            run_time=1.5
        )
        self.wait(0.5)
        
        # Create output layer
        self.play(
            Create(output_nodes),
            Create(output_label),
            run_time=1.5
        )
        self.wait(0.5)
        
        # Draw connections from input to hidden
        self.play(
            Create(input_to_hidden_connections),
            run_time=2
        )
        self.wait(0.5)
        
        # Draw connections from hidden to output
        self.play(
            Create(hidden_to_output_connections),
            run_time=2
        )
        self.wait(0.5)
        
        # Transform to show explanation
        self.play(
            Transform(
                VGroup(input_nodes, hidden_nodes, output_nodes, 
                       input_to_hidden_connections, hidden_to_output_connections,
                       input_label, hidden_label, output_label),
                VGroup(input_nodes, hidden_nodes, output_nodes, 
                       input_to_hidden_connections, hidden_to_output_connections,
                       input_label, hidden_label, output_label).copy().scale(0.85).shift(UP * 0.3)
            ),
            Create(explanation),
            run_time=2
        )
        
        self.wait(3)
`,
			Description: "Shows the basic structure of a neural network",
			SceneTitle:  "The Anatomy of a Neural Network",
		},
		{
			ClassName: "S03_HowInformationFlows",
			Code: `
from manim import *

class S03_HowInformationFlows(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#FFF9C4"
        
        # Create title
        title = Text(
            "How Information Flows",
            font_size=48,
            color="#424242",
            weight=BOLD
        ).shift(UP * 3.2)
        
        # Create input nodes (3 small circles)
        input_nodes = VGroup(*[
            Circle(
                radius=0.25,
                color="#FBC02D",
                fill_opacity=0.8,
                stroke_width=2,
                stroke_color="#424242"
            )
            for _ in range(3)
        ]).arrange(DOWN, buff=0.8)
        input_nodes.shift(LEFT * 5)
        
        # Create single neuron in focus
        focus_neuron = Circle(
            radius=0.5,
            color="#FBC02D",
            fill_opacity=0.9,
            stroke_width=3,
            stroke_color="#424242"
        ).shift(ORIGIN)
        
        # Create output nodes (2 small circles)
        output_nodes = VGroup(*[
            Circle(
                radius=0.25,
                color="#FBC02D",
                fill_opacity=0.8,
                stroke_width=2,
                stroke_color="#424242"
            )
            for _ in range(2)
        ]).arrange(DOWN, buff=1)
        output_nodes.shift(RIGHT * 5)
        
        # Create input values
        input_values = VGroup(
            Text("0.5", font_size=20, color="#424242"),
            Text("0.8", font_size=20, color="#424242"),
            Text("0.3", font_size=20, color="#424242")
        )
        for i, val in enumerate(input_values):
            val.move_to(input_nodes[i].get_center())
        
        # Create weighted connections with labels
        weights = [0.7, -0.4, 0.9]
        weight_labels = VGroup()
        connections = VGroup()
        
        for i, (node, weight) in enumerate(zip(input_nodes, weights)):
            connection = Line(
                node.get_center(),
                focus_neuron.get_center(),
                color="#FBC02D",
                stroke_width=2
            )
            connections.add(connection)
            
            weight_label = Text(
                f"w{i+1}={weight}",
                font_size=18,
                color="#424242"
            ).next_to(connection.get_center(), UP, buff=0.1)
            weight_labels.add(weight_label)
        
        # Create bias term
        bias_label = Text(
            "bias = 0.2",
            font_size=20,
            color="#424242",
            weight=BOLD
        ).next_to(focus_neuron, DOWN, buff=0.3)
        
        # Create weighted sum calculation
        calculation = VGroup(
            Text("Weighted Sum:", font_size=24, color="#424242", weight=BOLD),
            Text("0.5×0.7 + 0.8×(-0.4) + 0.3×0.9 + 0.2", font_size=20, color="#424242"),
            Text("= 0.35 - 0.32 + 0.27 + 0.2", font_size=20, color="#424242"),
            Text("= 0.5", font_size=20, color="#424242", weight=BOLD)
        ).arrange(DOWN, buff=0.2, aligned_edge=LEFT)
        calculation.shift(DOWN * 2.5)
        
        # Create activation function graphic
        axes = Axes(
            x_range=[-3, 3, 1],
            y_range=[0, 1, 0.5],
            x_length=3,
            y_length=1.5,
            axis_config={"color": "#424242"},
        ).shift(RIGHT * 2 + UP * 1.5)
        
        sigmoid = axes.plot(
            lambda x: 1 / (1 + np.exp(-x)),
            color="#FBC02D",
            stroke_width=3
        )
        
        activation_label = Text(
            "Activation Function",
            font_size=16,
            color="#424242"
        ).next_to(axes, UP, buff=0.2)
        
        # Create output value
        output_value = Text(
            "Output: 0.62",
            font_size=24,
            color="#424242",
            weight=BOLD
        ).next_to(focus_neuron, RIGHT, buff=1)
        
        # Create output connections
        output_connections = VGroup()
        for node in output_nodes:
            connection = Line(
                focus_neuron.get_center(),
                node.get_center(),
                color="#FBC02D",
                stroke_width=2
            )
            output_connections.add(connection)
        
        # Create labels
        weights_text = Text("Weights", font_size=20, color="#424242", weight=BOLD).shift(LEFT * 3 + UP * 2)
        bias_text = Text("Bias", font_size=20, color="#424242", weight=BOLD).next_to(bias_label, LEFT, buff=0.3)
        
        # Create explanation text
        explanation = Text(
            "Weights and biases are adjusted during training",
            font_size=22,
            color="#424242"
        ).shift(DOWN * 3.5)
        
        # Animations
        self.play(Write(title), run_time=1.5)
        self.wait(0.5)
        
        # Show input nodes and focus neuron
        self.play(
            Write(input_nodes),
            Write(focus_neuron),
            run_time=1.5
        )
        
        # Show input values
        self.play(Write(input_values), run_time=1)
        
        # Show weighted connections
        self.play(
            Write(connections),
            Write(weight_labels),
            Write(weights_text),
            run_time=2
        )
        self.wait(0.5)
        
        # Show bias
        self.play(
            Write(bias_label),
            Write(bias_text),
            run_time=1
        )
        
        # Show calculation
        self.play(Write(calculation), run_time=2)
        self.wait(1)
        
        # Show activation function
        self.play(
            Write(axes),
            Write(sigmoid),
            Write(activation_label),
            run_time=2
        )
        
        # Show output
        self.play(
            Write(output_value),
            Transform(
                calculation[-1].copy(),
                output_value
            ),
            run_time=1.5
        )
        
        # Show output connections
        self.play(
            Write(output_nodes),
            Write(output_connections),
            run_time=1.5
        )
        
        # Show training explanation
        self.play(Write(explanation), run_time=1.5)
        
        # Demonstrate weight adjustment
        new_weights = [0.75, -0.35, 0.85]
        new_weight_labels = VGroup()
        for i, (label, new_weight) in enumerate(zip(weight_labels, new_weights)):
            new_label = Text(
                f"w{i+1}={new_weight}",
                font_size=18,
                color="#424242"
            ).move_to(label.get_center())
            new_weight_labels.add(new_label)
        
        self.play(
            Transform(weight_labels, new_weight_labels),
            run_time=1.5
        )
        
        self.wait(2)
`,
			Description: "Explains how data moves through the network",
			SceneTitle:  "How Information Flows",
		},
		{
			ClassName: "S04_HowNeuralNetworksLearn",
			Code: `from manim import *

class S04_HowNeuralNetworksLearnTraining(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#E8F5E9"
        
        # Title
        title = Text(
            "How Neural Networks Learn (Training)",
            font_size=42,
            color="#1B5E20",
            weight=BOLD
        ).shift(UP * 3.3)
        
        # Create handwritten digit '7'
        digit_7 = VGroup()
        # Create a stylized 7
        seven_path = VMobject()
        seven_path.set_points_as_corners([
            [-0.4, 0.8, 0], [0.4, 0.8, 0], [0.2, 0.4, 0], 
            [-0.2, -0.4, 0], [-0.4, -0.8, 0]
        ])
        seven_path.set_stroke(color="#1B5E20", width=8)
        
        # Input frame
        input_frame = Square(
            side_length=1.8,
            color="#388E3C",
            fill_opacity=0.1,
            stroke_width=3
        )
        digit_7.add(input_frame, seven_path)
        digit_7.shift(LEFT * 5.5 + UP * 0.5)
        
        # Network structure
        # Input layer (3 nodes)
        input_layer = VGroup(*[
            Circle(radius=0.2, color="#388E3C", fill_opacity=0.8, stroke_width=2, stroke_color="#1B5E20")
            for _ in range(3)
        ]).arrange(DOWN, buff=0.5)
        input_layer.shift(LEFT * 3 + UP * 0.5)
        
        # Hidden layer (4 nodes)
        hidden_layer = VGroup(*[
            Circle(radius=0.2, color="#388E3C", fill_opacity=0.8, stroke_width=2, stroke_color="#1B5E20")
            for _ in range(4)
        ]).arrange(DOWN, buff=0.4)
        hidden_layer.shift(ORIGIN + UP * 0.5)
        
        # Output layer (10 nodes for digits 0-9)
        output_layer = VGroup(*[
            Circle(radius=0.15, color="#388E3C", fill_opacity=0.8, stroke_width=2, stroke_color="#1B5E20")
            for _ in range(10)
        ]).arrange(DOWN, buff=0.25)
        output_layer.shift(RIGHT * 3 + UP * 0.5)
        
        # Output labels
        output_labels = VGroup(*[
            Text(str(i), font_size=12, color="#1B5E20").move_to(output_layer[i])
            for i in range(10)
        ])
        
        # Connections
        connections_1 = VGroup()
        for i_node in input_layer:
            for h_node in hidden_layer:
                line = Line(
                    i_node.get_center(), h_node.get_center(),
                    color="#388E3C", stroke_width=1, stroke_opacity=0.5
                )
                connections_1.add(line)
        
        connections_2 = VGroup()
        for h_node in hidden_layer:
            for o_node in output_layer:
                line = Line(
                    h_node.get_center(), o_node.get_center(),
                    color="#388E3C", stroke_width=1, stroke_opacity=0.5
                )
                connections_2.add(line)
        
        # Network group
        network = VGroup(input_layer, hidden_layer, output_layer, connections_1, connections_2)
        
        # Output guess indicator
        guess_arrow = Arrow(
            start=output_layer[1].get_right() + RIGHT * 0.2,
            end=output_layer[1].get_right() + RIGHT * 1.2,
            color="#388E3C",
            stroke_width=3
        )
        guess_text = Text("Guess: 1", font_size=24, color="#1B5E20", weight=BOLD).next_to(guess_arrow, RIGHT)
        
        # Correct answer
        correct_text = Text("Correct: 7", font_size=24, color="#1B5E20", weight=BOLD).shift(RIGHT * 5 + DOWN * 1.5)
        
        # Error calculation
        error_box = Rectangle(
            width=3, height=0.8,
            color="#388E3C", fill_opacity=0.1, stroke_width=2
        ).shift(DOWN * 2.5)
        error_text = Text("Error = |1 - 7| = 6", font_size=20, color="#1B5E20").move_to(error_box)
        
        # Backpropagation arrow
        backprop_arrow = CurvedArrow(
            start_point=error_box.get_top(),
            end_point=network.get_bottom() + DOWN * 0.2,
            angle=-TAU/4,
            color="#388E3C",
            stroke_width=4
        )
        backprop_label = Text("Backpropagation", font_size=18, color="#1B5E20", weight=BOLD)
        backprop_label.next_to(backprop_arrow, LEFT, buff=0.3)
        
        # Training text
        training_text = Text("Training", font_size=20, color="#1B5E20", weight=BOLD).shift(LEFT * 5.5 + DOWN * 1.5)
        
        # Error reduction text
        error_reduction_text = Text("Error Reduction", font_size=20, color="#1B5E20", weight=BOLD).shift(DOWN * 3.5)
        
        # Animation sequence
        self.play(FadeIn(title), run_time=1)
        self.wait(0.5)
        
        # Show digit and network
        self.play(
            FadeIn(digit_7),
            FadeIn(network),
            FadeIn(output_labels),
            FadeIn(training_text),
            run_time=2
        )
        
        # First guess
        self.play(
            Indicate(output_layer[1], color="#388E3C", scale_factor=1.5),
            FadeIn(guess_arrow),
            FadeIn(guess_text),
            run_time=1.5
        )
        
        self.play(FadeIn(correct_text), run_time=1)
        
        # Show error
        self.play(
            FadeIn(error_box),
            FadeIn(error_text),
            run_time=1
        )
        
        # Backpropagation
        self.play(
            FadeIn(backprop_arrow),
            FadeIn(backprop_label),
            run_time=1.5
        )
        
        # Animate weight adjustment
        self.play(
            *[Indicate(line, color="#388E3C", scale_factor=1.02) for line in connections_1[::3]],
            *[Indicate(line, color="#388E3C", scale_factor=1.02) for line in connections_2[::3]],
            run_time=1.5
        )
        
        self.play(FadeIn(error_reduction_text), run_time=1)
        
        # Second iteration - better guess
        new_guess_text = Text("Guess: 3", font_size=24, color="#1B5E20", weight=BOLD).move_to(guess_text)
        new_error_text = Text("Error = |3 - 7| = 4", font_size=20, color="#1B5E20").move_to(error_text)
        
        self.play(
            Transform(guess_arrow, Arrow(
                start=output_layer[3].get_right() + RIGHT * 0.2,
                end=output_layer[3].get_right() + RIGHT * 1.2,
                color="#388E3C",
                stroke_width=3
            )),
            Transform(guess_text, new_guess_text),
            Indicate(output_layer[3], color="#388E3C", scale_factor=1.5),
            run_time=1.5
        )
        
        self.play(Transform(error_text, new_error_text), run_time=1)
        
        # Third iteration - even better
        final_guess_text = Text("Guess: 7", font_size=24, color="#1B5E20", weight=BOLD).move_to(guess_text)
        final_error_text = Text("Error = |7 - 7| = 0 ✓", font_size=20, color="#1B5E20").move_to(error_text)
        
        self.play(
            Transform(guess_arrow, Arrow(
                start=output_layer[7].get_right() + RIGHT * 0.2,
                end=output_layer[7].get_right() + RIGHT * 1.2,
                color="#388E3C",
                stroke_width=3
            )),
            Transform(guess_text, final_guess_text),
            Indicate(output_layer[7], color="#388E3C", scale_factor=1.5),
            run_time=1.5
        )
        
        self.play(
            Transform(error_text, final_error_text),
            Indicate(error_box, color="#388E3C", scale_factor=1.1),
            run_time=1.5
        )
        
        self.wait(2)
`,
			Description: "Visualizes the training process and backpropagation",
			SceneTitle:  "How Neural Networks Learn (Training)",
		},
		{
			ClassName: "S05_RealWorldApplications",
			Code: `from manim import *

class S05_RealWorldApplications(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#FCE4EC"
        
        # Title
        title = Text(
            "Real-World Applications",
            font_size=48,
            color="#880E4F",
            weight=BOLD
        ).shift(UP * 3.2)
        
        # Create smartphone icon for voice assistant
        smartphone = VGroup()
        phone_body = RoundedRectangle(
            width=1.2,
            height=2,
            corner_radius=0.2,
            color="#E91E63",
            fill_opacity=0.8,
            stroke_width=3
        )
        phone_screen = Rectangle(
            width=1,
            height=1.6,
            color="#880E4F",
            fill_opacity=0.3,
            stroke_width=2
        ).move_to(phone_body.get_center() + UP * 0.1)
        phone_button = Circle(
            radius=0.08,
            color="#880E4F",
            fill_opacity=1
        ).move_to(phone_body.get_bottom() + UP * 0.15)
        # Sound waves
        wave1 = Arc(radius=0.3, angle=PI/3, color="#E91E63", stroke_width=2).shift(RIGHT * 0.7)
        wave2 = Arc(radius=0.5, angle=PI/3, color="#E91E63", stroke_width=2).shift(RIGHT * 0.7)
        wave3 = Arc(radius=0.7, angle=PI/3, color="#E91E63", stroke_width=2).shift(RIGHT * 0.7)
        smartphone.add(phone_body, phone_screen, phone_button, wave1, wave2, wave3)
        smartphone.scale(0.8).shift(LEFT * 4.5 + UP * 1)
        
        smartphone_text = Text(
            "Voice Assistants",
            font_size=20,
            color="#880E4F",
            weight=BOLD
        ).next_to(smartphone, DOWN, buff=0.3)
        
        smartphone_desc = Text(
            "Neural networks process\nvoice commands",
            font_size=14,
            color="#880E4F"
        ).next_to(smartphone_text, DOWN, buff=0.2)
        
        # Create self-driving car icon
        car = VGroup()
        car_body = Rectangle(
            width=2.5,
            height=1,
            color="#E91E63",
            fill_opacity=0.8,
            stroke_width=3
        )
        car_roof = Polygon(
            [-0.8, 0.5, 0], [0.8, 0.5, 0], [0.5, 1, 0], [-0.5, 1, 0],
            color="#E91E63",
            fill_opacity=0.8,
            stroke_width=3
        )
        wheel1 = Circle(radius=0.25, color="#880E4F", fill_opacity=1).shift(LEFT * 0.7 + DOWN * 0.5)
        wheel2 = Circle(radius=0.25, color="#880E4F", fill_opacity=1).shift(RIGHT * 0.7 + DOWN * 0.5)
        # Sensor waves
        sensor1 = Arc(start_angle=-PI/4, angle=PI/2, radius=0.8, color="#E91E63", stroke_width=2).shift(RIGHT * 1.25)
        sensor2 = Arc(start_angle=-PI/4, angle=PI/2, radius=1.1, color="#E91E63", stroke_width=2).shift(RIGHT * 1.25)
        car.add(car_body, car_roof, wheel1, wheel2, sensor1, sensor2)
        car.scale(0.6).shift(RIGHT * 1.5 + UP * 1)
        
        car_text = Text(
            "Self-Driving Vehicles",
            font_size=20,
            color="#880E4F",
            weight=BOLD
        ).next_to(car, DOWN, buff=0.3)
        
        car_desc = Text(
            "Neural networks analyze\nroad conditions",
            font_size=14,
            color="#880E4F"
        ).next_to(car_text, DOWN, buff=0.2)
        
        # Create shopping cart icon
        cart = VGroup()
        cart_basket = Polygon(
            [-0.6, 0, 0], [0.6, 0, 0], [0.5, -0.8, 0], [-0.5, -0.8, 0],
            color="#E91E63",
            fill_opacity=0.8,
            stroke_width=3
        )
        cart_handle = Line(
            start=[-0.6, 0, 0],
            end=[-0.9, 0.4, 0],
            color="#E91E63",
            stroke_width=3
        )
        cart_wheel1 = Circle(radius=0.1, color="#880E4F", fill_opacity=1).shift(LEFT * 0.3 + DOWN * 0.9)
        cart_wheel2 = Circle(radius=0.1, color="#880E4F", fill_opacity=1).shift(RIGHT * 0.3 + DOWN * 0.9)
        # Items in cart
        item1 = Rectangle(width=0.2, height=0.3, color="#880E4F", fill_opacity=0.6).shift(UP * 0.2 + LEFT * 0.2)
        item2 = Rectangle(width=0.2, height=0.3, color="#880E4F", fill_opacity=0.6).shift(UP * 0.2 + RIGHT * 0.2)
        item3 = Rectangle(width=0.2, height=0.3, color="#880E4F", fill_opacity=0.6).shift(DOWN * 0.2)
        cart.add(cart_basket, cart_handle, cart_wheel1, cart_wheel2, item1, item2, item3)
        cart.scale(0.8).shift(LEFT * 4.5 + DOWN * 1.5)
        
        cart_text = Text(
            "Recommendation Systems",
            font_size=20,
            color="#880E4F",
            weight=BOLD
        ).next_to(cart, DOWN, buff=0.3)
        
        cart_desc = Text(
            "Neural networks predict\nuser preferences",
            font_size=14,
            color="#880E4F"
        ).next_to(cart_text, DOWN, buff=0.2)
        
        # Create medical scan icon
        medical = VGroup()
        scan_frame = Rectangle(
            width=1.8,
            height=1.8,
            color="#E91E63",
            fill_opacity=0.2,
            stroke_width=3
        )
        # Brain scan representation
        brain_outline = Ellipse(
            width=1.2,
            height=1,
            color="#880E4F",
            fill_opacity=0.3,
            stroke_width=2
        )
        # Highlighted region
        highlight_region = Ellipse(
            width=0.4,
            height=0.3,
            color="#E91E63",
            fill_opacity=0.8,
            stroke_width=2
        ).shift(RIGHT * 0.2 + UP * 0.1)
        # Cross lines
        cross_v = Line(UP * 0.9, DOWN * 0.9, color="#E91E63", stroke_width=1, stroke_opacity=0.5)
        cross_h = Line(LEFT * 0.9, RIGHT * 0.9, color="#E91E63", stroke_width=1, stroke_opacity=0.5)
        medical.add(scan_frame, brain_outline, highlight_region, cross_v, cross_h)
        medical.scale(0.8).shift(RIGHT * 1.5 + DOWN * 1.5)
        
        medical_text = Text(
            "Medical Image Recognition",
            font_size=20,
            color="#880E4F",
            weight=BOLD
        ).next_to(medical, DOWN, buff=0.3)
        
        medical_desc = Text(
            "Neural networks detect\ndiseases in scans",
            font_size=14,
            color="#880E4F"
        ).next_to(medical_text, DOWN, buff=0.2)
        
        # Group all applications
        app1 = VGroup(smartphone, smartphone_text, smartphone_desc)
        app2 = VGroup(car, car_text, car_desc)
        app3 = VGroup(cart, cart_text, cart_desc)
        app4 = VGroup(medical, medical_text, medical_desc)
        
        # Animation sequence
        self.play(Write(title), run_time=2)
        self.wait(0.5)
        
        # Fade in applications one by one
        self.play(FadeIn(app1), run_time=1.5)
        self.wait(0.5)
        
        self.play(FadeIn(app2), run_time=1.5)
        self.wait(0.5)
        
        self.play(FadeIn(app3), run_time=1.5)
        self.wait(0.5)
        
        self.play(FadeIn(app4), run_time=1.5)
        self.wait(0.5)
        
        # Highlight each application
        self.play(Indicate(smartphone, color="#E91E63", scale_factor=1.1), run_time=1)
        self.play(Indicate(car, color="#E91E63", scale_factor=1.1), run_time=1)
        self.play(Indicate(cart, color="#E91E63", scale_factor=1.1), run_time=1)
        self.play(Indicate(medical, color="#E91E63", scale_factor=1.1), run_time=1)
        
        self.wait(2)
`,
			Description: "Shows real-world uses of neural networks",
			SceneTitle:  "Real-World Applications",
		},
		{
			ClassName: "S06_ConclusionAndNextSteps",
			Code: `from manim import *

class S06_ConclusionAndNextSteps(Scene):
    def construct(self):
        # Set background color
        self.camera.background_color = "#FFFFFF"
        
        # Title
        title = Text(
            "Conclusion",
            font_size=56,
            color="#000000",
            weight=BOLD
        ).shift(UP * 3.2)
        
        # Summary points
        summary_title = Text(
            "What We've Learned:",
            font_size=32,
            color="#000000",
            weight=BOLD
        ).shift(UP * 2)
        
        # Key point 1
        point1 = VGroup(
            Circle(radius=0.08, color="#4CAF50", fill_opacity=1).shift(LEFT * 4.5 + UP * 0.8),
            Text(
                "Neural Networks are computing systems inspired by the brain",
                font_size=22,
                color="#000000"
            ).shift(LEFT * 0.8 + UP * 0.8)
        )
        
        # Key point 2
        point2 = VGroup(
            Circle(radius=0.08, color="#4CAF50", fill_opacity=1).shift(LEFT * 4.5 + UP * 0.1),
            Text(
                "Basic Structure: Input Layer → Hidden Layers → Output Layer",
                font_size=22,
                color="#000000"
            ).shift(LEFT * 0.8 + UP * 0.1)
        )
        
        # Key point 3
        point3 = VGroup(
            Circle(radius=0.08, color="#4CAF50", fill_opacity=1).shift(LEFT * 4.5 + DOWN * 0.6),
            Text(
                "Components: Neurons, Weights, Biases, Activation Functions",
                font_size=22,
                color="#000000"
            ).shift(LEFT * 0.8 + DOWN * 0.6)
        )
        
        # Key point 4
        point4 = VGroup(
            Circle(radius=0.08, color="#4CAF50", fill_opacity=1).shift(LEFT * 4.5 + DOWN * 1.3),
            Text(
                "Learning: Training through Backpropagation and Error Reduction",
                font_size=22,
                color="#000000"
            ).shift(LEFT * 0.6 + DOWN * 1.3)
        )
        
        # Key point 5
        point5 = VGroup(
            Circle(radius=0.08, color="#4CAF50", fill_opacity=1).shift(LEFT * 4.5 + DOWN * 2),
            Text(
                "Applications: Voice Assistants, Self-Driving Cars, Recommendations",
                font_size=22,
                color="#000000"
            ).shift(LEFT * 0.4 + DOWN * 2)
        )
        
        # Friendly message
        friendly_message = Text(
            "Neural Networks are powerful tools!",
            font_size=36,
            color="#4CAF50",
            weight=BOLD
        ).shift(DOWN * 0.5)
        
        # Encouragement text
        encouragement = VGroup(
            Text(
                "Keep exploring and learning!",
                font_size=28,
                color="#000000"
            ).shift(DOWN * 1.5),
            Text(
                "The journey into AI has just begun...",
                font_size=24,
                color="#000000",
                slant=ITALIC
            ).shift(DOWN * 2.2)
        )
        
        # Create decorative neural network icon
        nn_icon = VGroup()
        # Input nodes
        in1 = Circle(radius=0.1, color="#4CAF50", fill_opacity=0.8).shift(LEFT * 1.5 + UP * 0.3)
        in2 = Circle(radius=0.1, color="#4CAF50", fill_opacity=0.8).shift(LEFT * 1.5 + DOWN * 0.3)
        # Hidden nodes
        h1 = Circle(radius=0.1, color="#4CAF50", fill_opacity=0.8).shift(UP * 0.3)
        h2 = Circle(radius=0.1, color="#4CAF50", fill_opacity=0.8).shift(DOWN * 0.3)
        # Output node
        out = Circle(radius=0.1, color="#4CAF50", fill_opacity=0.8).shift(RIGHT * 1.5)
        # Connections
        lines = VGroup(
            Line(in1.get_center(), h1.get_center(), stroke_width=1, color="#4CAF50"),
            Line(in1.get_center(), h2.get_center(), stroke_width=1, color="#4CAF50"),
            Line(in2.get_center(), h1.get_center(), stroke_width=1, color="#4CAF50"),
            Line(in2.get_center(), h2.get_center(), stroke_width=1, color="#4CAF50"),
            Line(h1.get_center(), out.get_center(), stroke_width=1, color="#4CAF50"),
            Line(h2.get_center(), out.get_center(), stroke_width=1, color="#4CAF50")
        )
        nn_icon.add(lines, in1, in2, h1, h2, out)
        nn_icon.scale(1.2).shift(RIGHT * 4 + UP * 1)
        
        # Animation sequence
        self.play(FadeIn(title), run_time=2)
        self.wait(0.5)
        
        self.play(FadeIn(summary_title), run_time=1.5)
        self.wait(0.3)
        
        # Fade in summary points sequentially
        self.play(FadeIn(point1), run_time=1)
        self.wait(0.3)
        self.play(FadeIn(point2), run_time=1)
        self.wait(0.3)
        self.play(FadeIn(point3), run_time=1)
        self.wait(0.3)
        self.play(FadeIn(point4), run_time=1)
        self.wait(0.3)
        self.play(FadeIn(point5), run_time=1)
        self.wait(1)
        
        # Transform scene to show friendly message
        self.play(
            FadeIn(nn_icon),
            VGroup(summary_title, point1, point2, point3, point4, point5).animate.fade(0.3),
            run_time=2
        )
        
        self.play(FadeIn(friendly_message), run_time=2)
        self.wait(1)
        
        self.play(FadeIn(encouragement), run_time=2)
        
        self.wait(3)
`,
			Description: "Wraps up the video and encourages further learning",
			SceneTitle:  "Conclusion & Next Steps",
		},
	}

	// for _, scene := range scenes {
	// 	res, _ := utils.GenerateCode(scene, scenes, previousCode)
	// 	fmt.Println(res.SceneTitle)
	// 	// save to file
	// 	utils.WriteToFile(res.ClassName, res.Code)
	// 	generatedScenesMetaData = append(generatedScenesMetaData, *res)
	// 	pCode, _ := utils.ReadFile(res.ClassName)
	// 	previousCode = *pCode
	// 	fmt.Println("🟢 Code Generated For ", res.SceneTitle)
	// }

	for _, sc := range generatedScenesMetaData {

		func() {
    fmt.Println("🟢 Compiling For ", sc.SceneTitle)
			_, err := utils.CompileFile(sc.ClassName)
			if err != nil {
				fmt.Println("Fixing 🔴", sc.ClassName)
				utils.FixCode(fmt.Sprintf("%+v", err), sc)
			}
		}()
	}


}
