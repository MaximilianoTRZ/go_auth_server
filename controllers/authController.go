package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Go Auth Server!")
}
func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	

	return c.JSON(data)
}
func Login(c *fiber.Ctx) error {
	return c.SendString("Here will be the login endpoint")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Here will be the logout endpoint")
}