package controller

import (
	"time"

	"github.com/devianwahyu/farmigo/database"
	"github.com/devianwahyu/farmigo/model/entity"
	"github.com/devianwahyu/farmigo/model/request"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthRegister(c *fiber.Ctx) error {
	regData := new(request.RegisterRequest)

	// Get user request & check if error
	if err := c.BodyParser(regData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	// Validate user request
	if errors := request.ValidateRegisterStruct(*regData); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": errors,
		})
	}

	var user entity.User

	// Check email registered
	if database.DB.Where("email = ?", regData.Email).First(&user).RowsAffected > 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "failed",
			"message": "Email already taken by another user",
		})
	}

	// Hashing password
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to hash password",
		})
	}

	// Input user properties
	user.Name = regData.Name
	user.Email = regData.Email
	user.Password = string(hashedPassword)
	user.RoleID = regData.RoleID

	// Add new user to database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Email,
		"role": user.RoleID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"token":  t,
	})
}

func AuthLogin(c *fiber.Ctx) error {
	logData := new(request.LoginRequest)

	// Get user request & check if error
	if err := c.BodyParser(logData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	// Validate user request
	if errors := request.ValidateLoginStruct(*logData); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": errors,
		})
	}

	var user entity.User

	// Check account registered
	if database.DB.Where("email = ?", logData.Email).First(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "User not registered",
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logData.Password)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Password is wrong",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Email,
		"role": user.RoleID,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": t,
	})
}

func AuthChangePassword(c *fiber.Ctx) error {
	changePassData := new(request.ChangePasswordRequest)

	// Get user request & check if error
	if err := c.BodyParser(changePassData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	// Validate user request
	if errors := request.ValidateChangePasswordStruct(*changePassData); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": errors,
		})
	}

	var user entity.User

	// Check account registered
	if database.DB.Where("email = ?", changePassData.Email).First(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "User not registered",
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassData.OldPassword)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Password is wrong",
		})
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(changePassData.NewPassword), bcrypt.DefaultCost)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": errHash.Error(),
		})
	}

	// Update user password
	resultUdate := database.DB.Model(&user).Where("email = ?", changePassData.Email).Update("password", hashedPassword)
	if resultUdate.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": resultUdate.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Password updated successfuly",
	})
}
