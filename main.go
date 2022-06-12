package main

import (
	"github.com/Aviator-Coding/HttpPLC/middleware"
	"github.com/Aviator-Coding/HttpPLC/routes"
	"github.com/Aviator-Coding/HttpPLC/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init Fiber
	app := fiber.New()

	// Middleware
	middleware.FiberMiddleware(app)

	//routes
	routes.UserPublicRoute(app)
	routes.PlcPublicRoute(app)
	routes.SwaggerRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	utils.StartServer(app)
}
