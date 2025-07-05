package handler

import (
	"fmt"
	"github.com/bikaxh/vid-gen/primary-be/pkg/model"
	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateProjectHandler(c *fiber.Ctx) error {

	// get the body
	var project model.Project
	userId := c.Locals("userId").(string)
	err := c.BodyParser(&project)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid req body"})
	}

	project.UserId = userId
	// return plan
	text, _ := utils.GeneratePlan(project.Prompt)

	p, err := project.CreateProject(text)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project planned successfully",
		"plan":    text,
		"data":    p,
	})

}

func GenerateScenesHandler(c *fiber.Ctx) error {

	var project model.Project
	err := c.BodyParser(&project)

	if err != nil {
		fmt.Println("🔴", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}
	// save to db

	p, err := project.SavePlan()

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// push to queue

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Generating scenes",
		"data":    p,
	})
}

func CreateSceneHandler(c *fiber.Ctx) error {

	projectId := c.Params("projectId")

	var scene model.Scene
	err := c.BodyParser(&scene)
	
	fmt.Println("REqu body 🟢",scene)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid req body",
		})
	}

	s, err := scene.SaveScene(projectId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "something went wrong",
		})
	}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "scene created",
		"data":    s,
	})

}
