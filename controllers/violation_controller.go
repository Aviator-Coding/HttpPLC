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

// Register Violation
func CreateViolation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var Violation models.Violation

	//validate the request body
	if err := c.BodyParser(&Violation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&Violation); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create the New Violation
	CreateViolation := models.Violation{
		Text: Violation.Text,
	}

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Create(&CreateViolation)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.ViolationResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": CreateViolation}})
}

// Get a Violation By :violationId
func GetAViolationByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var Violation models.Violation
	defer cancel()

	log.Println(c.Params("violationId"))
	// Convert ViolationID to int
	violationId, err := strconv.ParseUint(c.Params("violationId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Violationid", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Where(&models.Violation{ID: uint(violationId)}).First(&Violation)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error - cannot update Database", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.ViolationResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": Violation}})
}

// Edit a Violation by :violationId
func EditAViolation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var Violation models.Violation

	violationId, err := strconv.ParseUint(c.Params("violationId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Violationid", Data: &fiber.Map{"data": err}})
	}

	//Parse the request body
	if err := c.BodyParser(&Violation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&Violation); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create the New Violation
	UpdateViolation := models.Violation{
		Text: Violation.Text,
	}

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Model(models.Violation{ID: uint(violationId)}).Updates(UpdateViolation)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error cannot writer to Database", Data: &fiber.Map{"data": result.Error.Error()}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.ViolationResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": UpdateViolation}})
}

// Delete a Violation by :violationId
func DeleteAViolation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//var Violation models.Violation
	defer cancel()

	// Convert ViolationID to int
	violationId, err := strconv.ParseUint(c.Params("violationId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Delete(&models.Violation{}, violationId)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.ViolationResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": violationId}})
}

// Get all Violations from the Database
func GetAllViolations(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var Violation models.Violation
	defer cancel()

	// Query all Violations From Database
	result := database.DB.WithContext(ctx).Find(&Violation)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ViolationResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.ViolationResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": Violation}})
}
