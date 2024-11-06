package main

import (
	"NyumbaUsersService/app/routes"
	"NyumbaUsersService/config"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.LoadEnv()

	// setting up a new fiber app router
	router := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "NyumbaUsersService v1.0.1",
		JSONEncoder:   json.Marshal,   // specifying the json encoder better performance
		JSONDecoder:   json.Unmarshal, // specifying the json decoder for better performance
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "*", //https://gofiber.io, https://gofiber.net,
		AllowHeaders:     "Origin, Content-Type, Accept, User-Agent, Content-Encoding",
		AllowCredentials: false,
		AllowMethods:     "*",
	}))

	// router.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Lax",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUIDv4,
	// }))

	routes.SetupRoutes(router)

	router.Listen(":3000")
}
