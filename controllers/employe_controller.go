package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Aviator-Coding/HttpPLC/database"
	"github.com/Aviator-Coding/HttpPLC/models"
	"github.com/Aviator-Coding/HttpPLC/responses"
	"github.com/gofiber/fiber/v2"
)

// Register Employe
func CreateEmploye(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var employe models.Employe

	//validate the request body
	if err := c.BodyParser(&employe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&employe); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create the New Employe
	// CreateEmploye := models.Employe{
	// 	FirstName:   employe.FirstName,
	// 	LastName:    employe.LastName,
	// 	Description: employe.Description,
	// 	BatchID:     employe.BatchID,
	// }

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Create(&employe)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.EmployeResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": employe}})
}

// Get a Employe By :employeid
func GetAEmployeByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var employe models.Employe
	defer cancel()

	log.Println(c.Params("employeId"))
	// Convert EmployeID to int
	employeId, err := strconv.ParseUint(c.Params("employeId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Employeid", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Where(&models.Employe{ID: uint(employeId)}).First(&employe)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error - cannot update Database", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.EmployeResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": employe}})
}

// Edit a Employe by :employeid
func EditAEmploye(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var employe models.Employe

	employeId, err := strconv.ParseUint(c.Params("employeId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Employeid", Data: &fiber.Map{"data": err}})
	}

	//Parse the request body
	if err := c.BodyParser(&employe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&employe); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create the New Employe
	// UpdateEmploye := models.Employe{
	// 	Name:     employe.Name,
	// 	Email:    employe.Email,
	// 	Password: password,
	// }

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Model(models.Employe{ID: uint(employeId)}).Updates(employe)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error cannot writer to Database", Data: &fiber.Map{"data": result.Error.Error()}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.EmployeResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": employe}})
}

// Delete a Employe by :employeid
func DeleteAEmploye(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//var employe models.Employe
	defer cancel()

	// Convert EmployeID to int
	employeId, err := strconv.ParseUint(c.Params("employeId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Delete(&models.Employe{}, employeId)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.EmployeResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": employeId}})
}

// Get all Employes from the Database
func GetAllEmployes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var employe models.Employe
	defer cancel()

	// Query all Employes From Database
	result := database.DB.WithContext(ctx).Find(&employe)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.EmployeResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.EmployeResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": employe}})
}
