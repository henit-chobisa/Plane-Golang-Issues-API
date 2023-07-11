package api

import (
	"fmt"
	"log"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start(host string, port int) {
	app := fiber.New(fiber.Config{
		AppName:          "PlaneAPI v1",
		Immutable:        false,
		JSONEncoder:      json.Marshal,
		JSONDecoder:      json.Unmarshal,
		DisableKeepalive: true,
	})

	// Apply Response Cache
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// Allow Cross-Origin Requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Initialize Request Logger
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Plane 👨🏼‍✈️")
	})

	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", host, port)))
}
