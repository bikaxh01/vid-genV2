package handler

import (
	"github.com/bikaxh/vid-gen/primary-be/pkg/model"
	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(c *fiber.Ctx) error {

	var user model.User

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request Body"})
	}

	existingUser, err := model.FindUserByEmail(user.Email)

	if existingUser != nil {

		return c.Status(fiber.StatusIMUsed).JSON(fiber.Map{"message": "Email already in use", "data": existingUser})
	}

	res, _ := user.Save()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user Created successfully", "data": res})

}

func SignInHandler(c *fiber.Ctx) error {

	var user model.User
	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Request Body"})
	}

	existingUser, err := model.FindUserByEmail(user.Email)

	if existingUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// match PW
	_, err = utils.ComparePassword(existingUser.Password, user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect password"})
	}

	// generate token
	token, err := utils.GenerateToken(existingUser.Email, existingUser.ID)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Something went wrong"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    *token,
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully logged in"})

}

func ValidateEmailHandler(c *fiber.Ctx) error {

	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	getUser, _ := model.FindUserByEmail(email)

	if getUser != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "email already in use",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Valid email 😁",
	})

}

func GetUserHandler(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user, err := model.FindUserById(userId)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "got user",
		"data":    user,
	})
}
