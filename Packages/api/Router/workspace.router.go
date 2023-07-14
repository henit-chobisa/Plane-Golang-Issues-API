package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func Initialize(group fiber.Router) {
	workspaceController := fiber.New()
	addWorkspaceHandlers(workspaceController)
	group.Mount("/workspaces", workspaceController)
}

func addWorkspaceHandlers(workspaceController *fiber.App) {
	workspaceController.Post("/create", func(c *fiber.Ctx) error {
		var workspace db.CreateWorkspaceParams
		if err := c.BodyParser(&workspace); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse JSON input, please check sent data.",
			})
		}

		if _, err := db.Db.CreateWorkspace(context.Background(), workspace); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create workspace",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(workspace)
	})

	workspaceController.Get("/get", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}
		workspace, err := db.Db.GetWorkspace(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve workspace",
			})
		}

		return c.JSON(workspace)
	})
}
