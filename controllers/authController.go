package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Go Auth Server!")
}
func Register(c *fiber.Ctx) error {
	return c.SendString("Here will be the registration endpoint")
}
func Login(c *fiber.Ctx) error {
	return c.SendString("Here will be the login endpoint")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Here will be the logout endpoint")
}