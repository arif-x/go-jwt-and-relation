package controllers

import (
	"errors"
	"strconv"
	"time"

	"go-relation/relasi-gorm/config"
	"go-relation/relasi-gorm/database"
	"go-relation/relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindByCredentials(email string, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	user := new(models.User)

	errFirst := database.DB.Where("email = ?", email).First(&user).Error

	if errFirst != nil {
		if errFirst == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("user not found")
	}

	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if check != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials
	user, err := FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    strconv.FormatInt(int64(user.ID), 10),
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	email := claims["email"].(string)
	return c.SendString("Welcome 👋" + email + " user id " + id)
}
