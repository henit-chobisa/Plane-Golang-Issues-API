package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func InitializeStatesRouter(group fiber.Router) {
	statesController := fiber.New()
	addStatesHandlers(statesController)
	group.Mount("/states", statesController)
}

func addStatesHandlers(statesController *fiber.App) {
	statesController.Post("/create", func(c *fiber.Ctx) error {
		var state db.CreateStateParams
		if err := c.BodyParser(&state); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		if _, err := db.Db.CreateState(context.Background(), state); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create state",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(state)
	})

	statesController.Get("/get/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}
		state, err := db.Db.GetState(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve project",
			})
		}

		return c.JSON(state)
	})
}
