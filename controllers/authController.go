package controllers

import (
	"go_react_auth/database"
	"go_react_auth/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Go Auth Server!")
}
func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	// send user to database as a reference
	database.DB.Create(&user)

	return c.JSON(user)
}
func Login(c *fiber.Ctx) error {
	return c.SendString("Here will be the login endpoint")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Here will be the logout endpoint")
}