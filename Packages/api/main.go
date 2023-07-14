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
	router "github.com/henit-chobisa/Plane-Golang-Issues-API/Packages/api/Router"
	db "github.com/henit-chobisa/Plane-Golang-Issues-API/db/sqlc"
)

func Start(host string, port int) {
	app := fiber.New(fiber.Config{
		AppName:          "PlaneAPI v1",
		Immutable:        false,
		JSONEncoder:      json.Marshal,
		JSONDecoder:      json.Unmarshal,
		DisableKeepalive: true,
	})

	// Initialize Database, initializes a global db object, that can be used for db interactions
	db.Initialize()

	// Apply Response Cache
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	app.Use(func(c *fiber.Ctx) error {
		// Set a custom header on all responses:
		c.Accepts(fiber.MIMEApplicationJSONCharsetUTF8)

		// Go to next middleware:
		return c.Next()
	})

	// Allow Cross-Origin Requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Initialize Request Logger
	app.Use(logger.New())

	// Initialize Routes
	v1 := app.Group("/v1")
	router.InitializeWorkspaceRouter(v1)
	router.InitializeUserRouter(v1)
	router.InitializeProjectsRouter(v1)
	router.InitializeStatesRouter(v1)
	router.InitializeIssuesRouter(v1)

	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", host, port)))
}
