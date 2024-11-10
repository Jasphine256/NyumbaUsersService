package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateSeeker(context *fiber.Ctx) error {
	user := new(structures.Seeker)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "invalid fields provided"})
	}

	success, message := services.AddSeekerToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(201).JSON(fiber.Map{"message": message})
}

func GetSeeker(context *fiber.Ctx) error {
	success, data := services.GetSeekerByEmail(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func UpdateSeeker(context *fiber.Ctx) error {
	success, data := services.UpdateSeekerById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func DeleteSeeker(context *fiber.Ctx) error {
	success, data := services.DeleteSeekerById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}
