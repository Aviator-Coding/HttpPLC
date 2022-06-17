package routes

import (
	"github.com/Aviator-Coding/HttpPLC/auth"
	"github.com/Aviator-Coding/HttpPLC/controllers"
	"github.com/gofiber/fiber/v2"
)

func EmployePublicRoute(app *fiber.App) {
	//All routes related to employes comes here
	app.Post("/employe", controllers.CreateEmploye)
	app.Get("/employe/:employeId", controllers.GetAEmployeByID)
	app.Put("/employe/:employeId", controllers.EditAEmploye)
	app.Delete("/employe/:employeId", controllers.DeleteAEmploye)
	app.Get("/employes", controllers.GetAllEmployes)
}

func EmployePrivatRoute(app *fiber.App) {
	app.Get("/employesecret", auth.AuthHandler, func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "HELLO FROM THE RESTRICTED AREA"})
	})
}
