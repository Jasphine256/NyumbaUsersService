package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateOwner(context *fiber.Ctx) error {
	user := new(structures.Owner)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": err})
	}

	success, message := services.AddOwnerToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(400).JSON(fiber.Map{"message": message})
}

func GetOwner(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func GetOwners() (bool, string) {
	success := true
	return success, "users list here"
}

func UpdateOwner(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func DeleteOwner(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}
