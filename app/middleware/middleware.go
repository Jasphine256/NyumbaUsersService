package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func SetupMiddleware(router *fiber.App) {

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*", //https://gofiber.io, https://gofiber.net,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
	}))

	router.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,
	}))

	// any custom middleware
	router.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
}
