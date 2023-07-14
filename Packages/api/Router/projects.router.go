package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func InitializeProjectsRouter(group fiber.Router) {
	projectsController := fiber.New()
	addProjectHandlers(projectsController)
	group.Mount("/projects", projectsController)
}

func addProjectHandlers(projectsController *fiber.App) {
	projectsController.Post("/create", func(c *fiber.Ctx) error {
		var project db.CreateProjectParams
		if err := c.BodyParser(&project); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		if _, err := db.Db.CreateProject(context.Background(), project); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create project",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(project)
	})

	projectsController.Get("/get/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}
		project, err := db.Db.GetProject(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve project",
			})
		}

		return c.JSON(project)
	})
}
