package controllers

import (
	"context"
	"github.com/STEEDUj2kb/pkg/repository"
	"github.com/STEEDUj2kb/pkg/utils"
	"strconv"
	"time"

	"github.com/STEEDUj2kb/app/models"
	"github.com/STEEDUj2kb/platform/database"
	"github.com/STEEDUj2kb/platform/ent/user"
	"github.com/gofiber/fiber/v2"
)

/*
TODO: refactor jwt validation
*/

// CreateStudy func creates study.
// @Description Create a study.
// @Summary creates a study
// @Tags Study
// @Accept json
// @Produce json
// @Success 201 {object} models.Study
// @Security ApiKeyAuth
// @Router /v1/study [post]
func CreateStudy(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `study:create` from JWT data of current study.
	credential := claims.Credentials[repository.StudyCreateCredential]

	// Only user with `study:create` credential can create a new study.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// Create a new study struct.
	study := new(models.Study)

	// Checking received data from JSON body.
	if err := c.BodyParser(study); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// TODO: Create a new validator for a study model.

	// get user from jwt claims.UserID
	authUser, err := database.EntClient.User.
		Query().
		Where(user.UUID(claims.UserID)).
		Only(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// create a study with given content
	createdStudy, err := database.EntClient.Study.Create().
		SetPlanner(authUser).
		SetContent(study.Content).
		Save(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Apply data from createdUser to user model
	study.ApplyData(createdStudy)

	// Return status 201 Created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"study": study,
	})
}

// GetStudies func get studies.
// @Description Get studies of request user.
// @Summary Get studies
// @Tags Study
// @Accept json
// @Produce json
// @Success 200 {object} models.Study
// @Security ApiKeyAuth
// @Router /v1/studies [get]
func GetStudies(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `study:create` from JWT data of current study.
	credential := claims.Credentials[repository.StudyCreateCredential]

	// Only user with `study:create` credential can create a new study.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// get user from jwt claims.UserID
	authUser, err := database.EntClient.User.
		Query().
		Where(user.UUID(claims.UserID)).
		Only(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	studiesQuery, err := authUser.QueryStudies().All(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	studies := make([]*models.Study, len(studiesQuery))
	for i, v := range studiesQuery {
		// Create a new study struct.
		study := new(models.Study)
		study.ApplyData(v)
		studies[i] = study
	}

	// Return status 200 ok.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"studies": studies,
	})
}

// CreateWeeklyGoal func creates a weekly goal by given study id.
// @Description Creates a weekly goal of request user by given study id.
// @Summary creates a weekly goal
// @Tags Study
// @Accept json
// @Produce json
// @Success 201 {object} models.WeeklyGaol
// @Security ApiKeyAuth
// @Router /v1/studies/{study_id}/weekly-goal [post]
func CreateWeeklyGoal(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `study:create` from JWT data of current study.
	credential := claims.Credentials[repository.StudyCreateCredential]

	// Only user with `study:create` credential can create a new study.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// Create a new weekly goal struct.
	wGoal := new(models.WeeklyGaol)

	// Checking received data from JSON body.
	if err := c.BodyParser(wGoal); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// TODO: Create a new validator for a weekly goal model.

	// get Study by given study_id param
	studyIdParam := c.Params("study_id")
	studyId, _ := strconv.Atoi(studyIdParam)

	// create a weekly goal
	createdWGoal, err := database.EntClient.WeeklyGaol.Create().
		SetStudyID(studyId).
		SetGoal(wGoal.Goal).
		SetNth(wGoal.Nth).
		Save(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	// Apply data from createdUser to user model
	wGoal.ApplyData(createdWGoal)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":       false,
		"msg":         nil,
		"weekly_goal": wGoal,
	})

}

// GetWeeklyGoals func gets current user's weekly goals by given study id.
// @Description Get weekly goals of a user study.
// @Summary get weekly goals
// @Tags Study
// @Accept json
// @Produce json
// @Success 200 {object} models.WeeklyGaol
// @Security ApiKeyAuth
// @Router /v1/studies/{study_id}/weekly-goals [get]
func GetWeeklyGoals(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Set credential `study:create` from JWT data of current study.
	credential := claims.Credentials[repository.StudyCreateCredential]

	// Only user with `study:create` credential can create a new study.
	if !credential {
		// Return status 403 and permission denied error message.
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": true,
			"msg":   "permission denied, check credentials of your token",
		})
	}

	// get Study by given study_id param
	studyIdParam := c.Params("study_id")
	studyId, _ := strconv.Atoi(studyIdParam)

	studyModel, err := database.EntClient.Study.Get(context.Background(), studyId)

	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	wgoalsQuery, err := studyModel.QueryWgoals().All(context.Background())
	if err != nil {
		// Return, if some problem with save model.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	wgoals := make([]*models.WeeklyGaol, len(wgoalsQuery))
	for i, v := range wgoalsQuery {
		// Create a new study struct.
		wgaol := new(models.WeeklyGaol)
		wgaol.ApplyData(v)
		wgoals[i] = wgaol
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"weekly_goals": wgoals,
	})
}
