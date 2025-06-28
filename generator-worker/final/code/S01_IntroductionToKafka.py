from manim import *

class S01_IntroductionToKafka(Scene):
    def construct(self):
        self.camera.background_color = "#1E1E1E"

        title = Text("Apache Kafka: A Beginner's Guide", color="#FFFFFF").scale(1.2)
        subtitle = Text("Understanding the Fundamentals", color="#FFFFFF").scale(0.7)
        subtitle.next_to(title, DOWN, buff=0.5)

        self.play(Write(title))
        self.play(FadeIn(subtitle))
        self.wait(2)
