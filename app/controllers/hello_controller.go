package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetHllo func gets first name.
// @Description Get first name with last name.
// @Summary get first name
// @Tags Hello
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
	last_name := strings.ToUpper(c.Params("last_name"))

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":      false,
		"msg":        nil,
		"first_name": people[last_name],
	})
}
