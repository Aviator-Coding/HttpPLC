package routes

import (
	"github.com/Aviator-Coding/HttpPLC/controllers"
	"github.com/gofiber/fiber/v2"
)

// Public Routing
func PlcPublicRoute(app *fiber.App) {
	//All routes related to users comes here
	plc := app.Group("/plc")
	plc.Post("/HMILogin", controllers.CreateHMIUser)
	plc.Put("/HMILogin/Machine/:batchId", controllers.CreateHMIUserMachine)
	plc.Get("/HMILogin/:machineName/:stationName/:batchId", controllers.GetHMIUserPermission)
	// plc.Get("/user/:userId", controllers.GetAUser)
	// plc.Put("/user/:userId", controllers.EditAUser)
	// plc.Delete("/user/:userId", controllers.DeleteAUser)
	// plc.Get("/users", controllers.GetAllUsers)
}

// Protected Routing with JWT Token
func PlcPrivatRoute(app *fiber.App) {

}
