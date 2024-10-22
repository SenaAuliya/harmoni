package routes

import (
	"harmoni/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	app.Post("auth/login", auth.Login)
	app.Post("auth/register", auth.Register)

}
