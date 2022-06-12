package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//var hmiUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "HMIUsers")

// Creates an HMI User
func CreateHMIUser(c *fiber.Ctx) error {

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// var plcHMIUser models.PlcHMIUser

	// //validate the request body
	// if err := c.BodyParser(&plcHMIUser); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	// //use the validator library to validate required fields
	// if validationErr := validate.Struct(&plcHMIUser); validationErr != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	// }

	// // Insert Results into collection
	// result, err := hmiUserCollection.InsertOne(ctx, models.PlcHMIUser{
	// 	FirstName: plcHMIUser.FirstName,
	// 	LastName:  plcHMIUser.LastName,
	// 	BatchID:   plcHMIUser.BatchID,
	// 	CreatedAt: time.Now(),
	// 	Comment:   plcHMIUser.Comment,
	// })
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{Status: fiber.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }
	// return c.Status(fiber.StatusCreated).JSON(responses.UserResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
	return nil
}

func CreateHMIUserMachine(c *fiber.Ctx) error {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	// 	var plcMachines models.PLCMachines
	// 	//var plcHMIUser models.PlcHMIUser
	// 	batchId := c.Params("batchId")

	// 	//validate the request body
	// 	if err := c.BodyParser(&plcMachines); err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// 	}

	// 	//use the validator library to validate required fields
	// 	if validationErr := validate.Struct(&plcMachines); validationErr != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	// 	}

	// 	// newUser := models.User{
	// 	// 	Id:       primitive.NewObjectID(),
	// 	// 	Name:     user.Name,
	// 	// 	Location: user.Location,
	// 	// 	Title:    user.Title,
	// 	// }

	// 	update := bson.M{
	// 		"$addToSet": bson.M{
	// 			"machines": bson.M{"$each": []string{"camera", "electronics", "accessories"}},
	// 		},
	// 	}

	// 	result, err := hmiUserCollection.UpdateOne(ctx, bson.M{"batchid": batchId}, update)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{Status: fiber.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// 	}
	// 	return c.Status(fiber.StatusCreated).JSON(responses.UserResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
	return nil
}

// Get the HMI User Permission
func GetHMIUserPermission(c *fiber.Ctx) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// var plcHMIUser models.PlcHMIUser
	// batchID := c.Params("batchId")
	// // MachineName := c.Params("machineName")
	// // StationsName := c.Params("stationName")

	// // Find the User with the Permission
	// err := hmiUserCollection.FindOne(ctx, bson.M{"batchid": batchID}).Decode(&plcHMIUser)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{Status: fiber.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	// // If login was okay we need to Update it

	// return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": plcHMIUser}})
	return nil
}
