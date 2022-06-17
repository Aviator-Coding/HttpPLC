package main

import (
	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/Aviator-Coding/HttpPLC/database"
	"github.com/Aviator-Coding/HttpPLC/middleware"
	"github.com/Aviator-Coding/HttpPLC/routes"
	"github.com/Aviator-Coding/HttpPLC/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init Fiber
	app := fiber.New()

	//Load Env Files
	configs.LoadConfig()

	// Connect to DB
	database.ConnectDB()

	// Middleware
	middleware.FiberMiddleware(app)

	//Unrestricted Routes
	routes.UserPublicRoute(app)
	routes.EmployePublicRoute(app)
	routes.ViolationPublicRoute(app)
	routes.PlcPublicRoute(app)
	routes.SwaggerRoute(app)

	routes.Configure(app)

	// JWT Restricted Routes
	routes.UserPrivatRoute(app)
	routes.ViolationPrivatRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	routes.NotFoundRoute(app)
	utils.StartServer(app)
}
