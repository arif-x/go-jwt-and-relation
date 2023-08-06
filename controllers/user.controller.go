package controllers

import (
	"go-relation/relasi-gorm/database"
	"go-relation/relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Locker").Preload("Posts").Find(&users)

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.UserStore)

	// PARSE BODY REQUEST TO OBJECT STRUCT
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "can't handle request",
		})
	}

	// NAME MANUAL VALIDATION
	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	// EMAIL MANUAL VALIDATION
	if user.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "email is required",
		})
	}

	// PASSWORD MANUAL VALIDATION
	if user.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "password is required",
		})
	}

	password := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "create data successfully",
			"user":    nil,
		})
	}

	user.Password = string(hashedPassword)

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message": "create data successfully",
		"user":    user,
	})
}
