package main

import (
	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/Aviator-Coding/HttpPLC/middleware"
	"github.com/Aviator-Coding/HttpPLC/routes"
	"github.com/Aviator-Coding/HttpPLC/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()
	utils.CreateIndex()

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
