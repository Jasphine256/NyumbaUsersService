package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateOwner(context *fiber.Ctx) error {
	user := new(structures.Owner)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "invalid fields provided"})
	}

	success, message := services.AddOwnerToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(201).JSON(fiber.Map{"message": message})
}

func GetOwner(context *fiber.Ctx) error {
	success, data := services.GetOwnerByEmail(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func UpdateOwner(context *fiber.Ctx) error {
	success, message := services.UpdateOwnerById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}

	return context.Status(200).JSON(fiber.Map{"message": message})

}

func DeleteOwner(context *fiber.Ctx) error {
	success, data := services.DeleteOwnerById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}
