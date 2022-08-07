package router

import (
	"github.com/devianwahyu/farmigo/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r fiber.Router) {
	r.Post("/register", controller.AuthRegister)
	r.Post("/login", controller.AuthLogin)
	// r.Put("/change-password")
}
