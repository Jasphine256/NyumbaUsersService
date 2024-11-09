package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateTenant(context *fiber.Ctx) error {
	user := new(structures.Tenant)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": err})
	}

	success, message := services.AddTenantToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(400).JSON(fiber.Map{"message": message})
}

func GetTenant(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func GetTenants() (bool, string) {
	success := true
	return success, "users list here"
}

func UpdateTenant(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func DeleteTenant(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}
