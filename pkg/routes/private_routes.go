package routes

import (
	"github.com/STEEDUj2kb/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of public routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for study related
	route.Post("/study", controllers.CreateStudy)                              // create a study
	route.Get("/studies", controllers.GetStudies)                              // get all studies of the authenticated user
	route.Post("/studies/:study_id/weekly-goal", controllers.CreateWeeklyGoal) // create a weekly goal
	route.Get("/studies/:study_id/weekly-goals", controllers.GetWeeklyGoals)   // get weekly goals
}
