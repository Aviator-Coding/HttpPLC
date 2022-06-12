package main

import (
	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/Aviator-Coding/HttpPLC/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())

	//routes
	routes.UserPublicRoute(app) //add this
	routes.SwaggerRoute(app)    //add this

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	app.Listen(configs.EnvServerUrl())
}
