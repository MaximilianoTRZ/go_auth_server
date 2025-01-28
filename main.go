package main

import (
	"go_react_auth/database"
	"go_react_auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

    database.Connect()  

    app := fiber.New()

    app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

    routes.SetupRoutes(app)

    app.Listen(":8080")
}