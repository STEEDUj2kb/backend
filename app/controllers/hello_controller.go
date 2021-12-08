package controllers

import (
	"context"
	"strings"

	"github.com/STEEDUj2kb/v1/app/models"
	"github.com/STEEDUj2kb/v1/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetHello func gets first name.
// @Description Get first name with last name.
// @Summary get first name
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /v1/hello/{last_name} [get]
func GetHello(c *fiber.Ctx) error {
	people := map[string]string{
		"HONG": "GIL DONG",
		"KIM":  "GA NAE",
		"SHIN": "GI RU",
	}
	lastName := strings.ToUpper(c.Params("last_name"))

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":      false,
		"msg":        nil,
		"first_name": people[lastName],
	})
}

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	user := new(models.User)
	// Checking received data from JSON body.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	createdUser, err := database.EntClient.User.
		Create().
		SetEmail(user.Email).
		SetName(user.Name).
		Save(context.Background())

	if err != nil {
		c.Status(500).JSON("Unable to save note")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201 OK.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  createdUser,
	})

}
