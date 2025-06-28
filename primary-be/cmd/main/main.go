package main

import (
	"fmt"

	"github.com/bikaxh/vid-gen/primary-be/pkg/handler"
	"github.com/bikaxh/vid-gen/primary-be/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	fmt.Println("Main ")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))
	app.Get("/ping", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(map[string]string{"message": "PONG"})
	})

	userRouter := app.Group("/user")
	userRouter.Post("/sign-up", handler.SignUpHandler)
	userRouter.Post("/sign-in", handler.SignInHandler)
	userRouter.Get("/validate-email/:email", handler.ValidateEmailHandler)
	userRouter.Get("/me", middleware.AuthMiddleware(), handler.GetUserHandler)

	projectRouter := app.Group("/project", middleware.AuthMiddleware())
	projectRouter.Post("/create-project", handler.CreateProjectHandler)
	projectRouter.Post("/generate-scenes", handler.GenerateScenesHandler)

	app.Listen(":8000")

}
