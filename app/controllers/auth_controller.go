package controllers

import (
	"context"
	"strings"

	"github.com/STEEDUj2kb/v1/app/models"
	"github.com/STEEDUj2kb/v1/pkg/utils"
	"github.com/STEEDUj2kb/v1/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
// @Param name body string true "Name"
// @Param password body string true "Password"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func UserSignUp(c *fiber.Ctx) error {
	defer database.EntClient.Close()

	// Create a new user auth struct.
	signUp := new(models.SignUp)

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	user := &models.User{
		UUID:         uuid.New(),
		Name:         signUp.Name,
		Email:        signUp.Email,
		PasswordHash: utils.GeneratePassword(signUp.Password),
	}

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	createdUser, err := database.EntClient.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPasswordHash(user.PasswordHash).
		Save(context.Background())

	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Apply data from createdUser to user model
	user.ApplyData(createdUser)

	// Return status 201 OK.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})

}
