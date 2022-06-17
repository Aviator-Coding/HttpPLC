package routes

import (
	"github.com/Aviator-Coding/HttpPLC/controllers"
	"github.com/gofiber/fiber/v2"
)

func ViolationPublicRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/violation", controllers.CreateViolation)
	app.Get("/violation/:violationId", controllers.GetAViolationByID)
	app.Put("/violation/:violationId", controllers.EditAViolation)
	app.Delete("/violation/:violationId", controllers.DeleteAViolation)
	app.Get("/violation", controllers.GetAllViolations)
}

func ViolationPrivatRoute(app *fiber.App) {
	// app.Get("/usersecret", auth.AuthHandler, func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "HELLO FROM THE RESTRICTED AREA"})
	// })
}
