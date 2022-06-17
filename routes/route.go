package routes

import (
	"github.com/Aviator-Coding/HttpPLC/auth"
	"github.com/Aviator-Coding/HttpPLC/controllers"
	"github.com/gofiber/fiber/v2"
)

func Configure(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1") // /api/v1
	ConfigureV1(v1)        // /api/v1/list

}

func ConfigureV1(version fiber.Router) {
	//All routes related to users comes here
	version.Post("/user", controllers.CreateUser)
	version.Get("/user/:userId", controllers.GetAUserByID)
	version.Put("/user/:userId", controllers.EditAUser)
	version.Delete("/user/:userId", controllers.DeleteAUser)
	version.Get("/users", controllers.GetAllUsers)
	version.Post("/user/login", controllers.LoginUser)
	version.Post("/user/logout", controllers.LogoutUser)

	version.Get("/usersecret", auth.AuthHandler, func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "HELLO FROM THE RESTRICTED AREA"})
	})
}
