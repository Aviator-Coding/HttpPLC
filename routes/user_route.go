package routes

import (
	"github.com/Aviator-Coding/HttpPLC/auth"
	"github.com/Aviator-Coding/HttpPLC/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserPublicRoute(app *fiber.App) {
	//All routes related to users comes here
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUserByID)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Post("/user/login", controllers.LoginUser)
	app.Post("/user/logout", controllers.LogoutUser)
}

func UserPrivatRoute(app *fiber.App) {
	app.Get("/usersecret", auth.AuthHandler, func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "HELLO FROM THE RESTRICTED AREA"})
	})
}
