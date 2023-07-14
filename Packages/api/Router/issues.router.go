package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func InitializeIssuesRouter(group fiber.Router) {
	issuesController := fiber.New()
	addIssuesHandlers(issuesController)
	group.Mount("/issues", issuesController)
}

func addIssuesHandlers(issuesController *fiber.App) {
	issuesController.Post("/create", func(c *fiber.Ctx) error {
		var issue db.CreateIssueParams
		if err := c.BodyParser(&issue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		if _, err := db.Db.CreateIssue(context.Background(), issue); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create issue",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(issue)
	})

	issuesController.Get("/get/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}
		issue, err := db.Db.GetIssue(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve issues",
			})
		}

		return c.JSON(issue)
	})

	issuesController.Get("/getByProject/:projectid", func(c *fiber.Ctx) error {
		project_id := c.Params("projectid")
		uuid, err := uuid.Parse(project_id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}

		issues, err := db.Db.ListIssuesByProject(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve issues",
			})
		}

		return c.JSON(issues)
	})

}
