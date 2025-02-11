package routes

import (
	"go_react_auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/auth", controllers.Hello)
	app.Get("/auth/authenticated-user", controllers.User)
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", controllers.Logout)

}