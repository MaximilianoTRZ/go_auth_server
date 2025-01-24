package main

import (
	"go_react_auth/database"
	"go_react_auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

    database.Connect()  

    app := fiber.New()

    routes.SetupRoutes(app)

    app.Listen(":8080")
}