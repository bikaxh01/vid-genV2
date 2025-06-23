package middleware

import (
	"fmt"

	"strings"

	"github.com/bikaxh/vid-gen/primary-be/pkg/model"
	"github.com/bikaxh/vid-gen/primary-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {

		byteToken := c.Request().Header.Peek("Authorization")

		bearerToken := string(byteToken)

		if bearerToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized user",
			})
		}

		token := strings.Split(bearerToken, " ")[1]

		authData, err := utils.VerifyToken(token)

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized user",
			})
		}
		user, err := model.FindUserById(*authData)

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}

		c.Locals("userId", user.ID)

		return c.Next()
	}
}
