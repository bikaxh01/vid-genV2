from manim import *

class S02_TheProblemTraditionalMessaging(Scene):
    def construct(self):
        # Color scheme
        background_color = "#1E1E1E"
        highlight_color = "#00CED1"
        text_color = "#FFFFFF"

        # Text elements
        title = Text("Traditional Messaging", color=text_color).to_edge(UP)
        bottleneck_text = Text("Bottlenecks & Complexity", color=text_color).to_edge(DOWN)

        # Diagram elements
        producer = Circle(radius=0.5, color=highlight_color, fill_opacity=1).shift(LEFT * 4)
        producer_label = Text("Producer", color=text_color).next_to(producer, DOWN)
        consumer = Square(side_length=1, color=highlight_color, fill_opacity=1).shift(RIGHT * 4)
        consumer_label = Text("Consumer", color=text_color).next_to(consumer, DOWN)

        # Arrow (Slow Data Transfer)
        arrow = Arrow(producer.get_right(), consumer.get_left(), color=highlight_color, stroke_width=5)
        arrow_label = Text("Slow Data Transfer", color=text_color).move_to(arrow.get_center() + UP * 0.5).scale(0.7)

        # Indicate Bottleneck
        brace = Brace(arrow, direction=DOWN, color=highlight_color)

        # Animations
        self.play(Write(title))
        self.play(Create(producer), Write(producer_label))
        self.play(Create(consumer), Write(consumer_label))
        self.play(Create(arrow))
        self.play(Write(arrow_label))
        self.play(Create(brace))
        self.play(Write(bottleneck_text))

        self.wait(3)
        self.play(FadeOut(title, producer, producer_label, consumer, consumer_label, arrow, arrow_label, brace, bottleneck_text))
        self.wait(1)