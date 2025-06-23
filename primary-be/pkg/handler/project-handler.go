package handler

import (


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
	//save project with pending

	// return plan
	text, _ := utils.GeneratePlan(project.Prompt)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project planned successfully",
		"data":    text,
	})

}
