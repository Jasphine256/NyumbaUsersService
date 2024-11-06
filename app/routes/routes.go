package routes

import (
	"NyumbaUsersService/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	},
	)

	admins := router.Group("/admins")
	users := router.Group("/users")

	// #####################  ROUTES FOR /users/* #####################

	users.Route("/owners", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("hello")
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.SendString(id)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			success, data := controllers.CreateOwner(c)
			if !success {
				return c.Status(400).JSON(fiber.Map{"message": data})
			}
			return c.Status(201).JSON(fiber.Map{"message": "created", "data": data})
		})

		api.Patch("/:id/:data", func(c *fiber.Ctx) error {
			id := c.Params("id")
			data := c.Params("data")
			return c.JSON(fiber.Map{"id": id, "data": data})
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return c.SendString("deleted")
		})
	})

	users.Route("/seekers", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("hello")
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.SendString(id)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			success, data := controllers.CreateSeeker(c)
			if !success {
				return c.Status(400).JSON(fiber.Map{"message": data})
			}
			return c.Status(201).JSON(fiber.Map{"message": "created", "data": data})
		})

		api.Patch("/:id/:data", func(c *fiber.Ctx) error {
			id := c.Params("id")
			data := c.Params("data")
			return c.JSON(fiber.Map{"id": id, "data": data})
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return c.SendString("deleted")
		})
	})

	users.Route("/tenants", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("hello")
		})

		api.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.SendString(id)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			success, data := controllers.CreateTenant(c)
			if !success {
				return c.Status(400).JSON(fiber.Map{"message": data})
			}
			return c.Status(201).JSON(fiber.Map{"message": "created", "data": data})
		})

		api.Patch("/:id/:data", func(c *fiber.Ctx) error {
			id := c.Params("id")
			data := c.Params("data")
			return c.JSON(fiber.Map{"id": id, "data": data})
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return c.SendString("deleted")
		})
	})

	// #####################  ROUTES FOR /admins/* #####################

	// GET ALL ADMINS
	admins.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello admin")
	})

	// GET ADMIN BY ID
	admins.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(id)
	})

	// CREATE NEW ADMIN
	admins.Post("/new", func(c *fiber.Ctx) error {
		success, data := controllers.CreateAdmin(c)
		if !success {
			return c.Status(400).JSON(fiber.Map{"message": data})
		}
		return c.Status(201).JSON(fiber.Map{"message": "created", "data": data})
	})

	// UPDATE ADMIN DATA
	admins.Patch("/:id/:data", func(c *fiber.Ctx) error {
		id := c.Params("id")
		data := c.Params("data")
		return c.JSON(fiber.Map{"id": id, "data": data})
	})

	// DELETE ADMIN BY ID
	admins.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("deleted")
	})
}
