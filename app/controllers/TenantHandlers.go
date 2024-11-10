package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateTenant(context *fiber.Ctx) error {
	user := new(structures.Tenant)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "invalid fields provided"})
	}

	success, message := services.AddTenantToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(201).JSON(fiber.Map{"message": message})
}

func GetTenant(context *fiber.Ctx) error {
	success, data := services.GetTenantByEmail(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func UpdateTenant(context *fiber.Ctx) error {
	success, data := services.UpdateTenantById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func DeleteTenant(context *fiber.Ctx) error {
	success, data := services.DeleteTenantById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}
