package main

import (
	"fmt"

	"github.com/bikaxh/vid-gen/primary-be/pkg/handler"
	"github.com/bikaxh/vid-gen/primary-be/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Main ")

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(map[string]string{"message": "PONG"})
	})

	userRouter := app.Group("/user")
	userRouter.Post("/sign-up", handler.SignUpHandler)
	userRouter.Post("/sign-in",  handler.SignInHandler)

	projectRouter := app.Group("/project", middleware.AuthMiddleware())
	projectRouter.Post("/create-project", handler.CreateProjectHandler)

	app.Listen(":8000")

}
