package handler

import (
	"github.com/bikaxh/vid-gen/primary-be/pkg/model"
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
	
	// generate plan

	//save project with pending

	// return plan

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
		"data":userId,
	})

}
