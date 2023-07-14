package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func InitializeUserRouter(group fiber.Router) {
	userController := fiber.New()
	addUserHandlers(userController)
	group.Mount("/users", userController)
}

func addUserHandlers(userController *fiber.App) {
	userController.Post("/create", func(c *fiber.Ctx) error {
		var user db.CreateUserParams
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		if _, err := db.Db.CreateUser(context.Background(), user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})

	userController.Get("/get/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		uuid, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid UUID format",
			})
		}
		user, err := db.Db.GetUser(context.Background(), uuid)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve user",
			})
		}

		return c.JSON(user)
	})
}
