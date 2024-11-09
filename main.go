package main

import (
	"NyumbaUsersService/app/middleware"
	"NyumbaUsersService/app/routes"
	"NyumbaUsersService/config"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
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

	middleware.SetupMiddleware(router)
	middleware.SetupSecurity(router)

	routes.SetupRoutes(router)

	router.Listen(":3000")
}
