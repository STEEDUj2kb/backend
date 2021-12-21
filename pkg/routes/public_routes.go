package routes

import (
	"github.com/STEEDUj2kb/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for user
	route.Get("/hello/:last_name", controllers.GetHello) // get first_name by last_name
	route.Post("/user/sign/up", controllers.UserSignUp)  // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn)  // auth, return Access & Refresh tokens
}
