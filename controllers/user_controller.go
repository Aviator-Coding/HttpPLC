package controllers

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Aviator-Coding/HttpPLC/auth"
	"github.com/Aviator-Coding/HttpPLC/database"
	"github.com/Aviator-Coding/HttpPLC/models"
	"github.com/Aviator-Coding/HttpPLC/responses"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register User
func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create a Password and Validate
	password, bcryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcryptErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": bcryptErr.Error()}})
	}

	// Create the New User
	CreateUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Create(&CreateUser)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.UserResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": CreateUser}})
}

// Get a User By :userid
func GetAUserByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	log.Println(c.Params("userId"))
	// Convert UserID to int
	userId, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Userid", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Where(&models.User{ID: uint(userId)}).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error - cannot update Database", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// Edit a User by :userid
func EditAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User

	userId, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error - Cannot convert Userid", Data: &fiber.Map{"data": err}})
	}

	//Parse the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Create a Password and Validate
	password, bcryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if bcryptErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": bcryptErr.Error()}})
	}

	// Create the New User
	UpdateUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: password,
	}

	// Insert Data into Database
	result := database.DB.WithContext(ctx).Model(models.User{ID: uint(userId)}).Updates(UpdateUser)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error cannot writer to Database", Data: &fiber.Map{"data": result.Error.Error()}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusCreated).JSON(responses.UserResponse{Status: fiber.StatusCreated, Message: "success", Data: &fiber.Map{"data": UpdateUser}})
}

// Delete a User by :userid
func DeleteAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//var user models.User
	defer cancel()

	// Convert UserID to int
	userId, err := strconv.ParseUint(c.Params("userId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err}})
	}

	// Query Data From Database
	result := database.DB.WithContext(ctx).Delete(&models.User{}, userId)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": userId}})
}

// Get all Users from the Database
func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	// Query all Users From Database
	result := database.DB.WithContext(ctx).Find(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": result.Error}})
	}

	// Repsone to the Client that the Operation was succesfull
	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
}

// Login a User (S)
func LoginUser(c *fiber.Ctx) error {
	var userPost models.User
	var user models.User

	//Bodyparse the request body
	if err := c.BodyParser(&userPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Find the User with his email or Username address
	database.DB.Where("email = ?", userPost.Email).First(&user)
	if user.ID == 0 {
		database.DB.Where("email = ?", userPost.Name).First(&user)
		if user.ID == 0 {
			return c.Status(fiber.StatusNotFound).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "User not found", Data: &fiber.Map{"data": ""}})
		}
	}

	// Check user Password against the Databse
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(userPost.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "Password is not correct", Data: &fiber.Map{"data": ""}})
	}

	token, errLogin := auth.LoginUser(c, user)
	if errLogin != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.UserResponse{Status: fiber.StatusBadRequest, Message: "JWT Token Error", Data: &fiber.Map{"data": ""}})
	}
	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"token": token}})
}

// Logsout User
func LogoutUser(c *fiber.Ctx) error {
	auth.LogOutUser(c)
	return c.Status(fiber.StatusOK).JSON(responses.UserResponse{Status: fiber.StatusOK, Message: "success", Data: &fiber.Map{"data": "Successfull Logged out"}})
}
