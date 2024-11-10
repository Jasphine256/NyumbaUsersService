package routes

import (
	"NyumbaUsersService/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router *fiber.App) {

	// router.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("WELCOME TO JASPHINE DIGITAL TECHNOLOGIES")
	// },
	// )

	// #####################  ROUTER GROUPS #####################

	admins := router.Group("/admins")
	users := router.Group("/users")

	// #####################  ROUTES FOR /users/* #####################

	users.Route("/owners", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetOwner(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateOwner(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateOwner(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteOwner(c)
		})
	})

	users.Route("/seekers", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetSeeker(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateSeeker(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateSeeker(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteSeeker(c)
		})
	})

	users.Route("/tenants", func(api fiber.Router) {

		api.Get("/", func(c *fiber.Ctx) error {
			return controllers.GetTenant(c)
		})

		api.Post("/new", func(c *fiber.Ctx) error {
			return controllers.CreateTenant(c)
		})

		api.Patch("/:id", func(c *fiber.Ctx) error {
			return controllers.UpdateTenant(c)
		})

		api.Delete("/:id", func(c *fiber.Ctx) error {
			return controllers.DeleteTenant(c)
		})
	})

	// #####################  ROUTES FOR /admins/* #####################

	// CREATE NEW ADMIN
	admins.Post("/new", func(c *fiber.Ctx) error {
		return controllers.CreateAdmin(c)
	})

	// GET ADMIN
	admins.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetAdmin(c)
	})

	// UPDATE ADMIN DATA
	admins.Patch("/:id", func(c *fiber.Ctx) error {
		return controllers.UpdateAdmin(c)
	})

	// DELETE ADMIN BY ID
	admins.Delete("/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteAdmin(c)
	})
}
