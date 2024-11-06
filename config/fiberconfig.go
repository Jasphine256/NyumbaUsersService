package config

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouterConfig(router *fiber.App) {

	router.Static("/static", "./public")
	// => http://localhost:3000/static/hello.html looks in public folder

}

