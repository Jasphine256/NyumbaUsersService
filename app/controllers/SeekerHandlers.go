package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateSeeker(context *fiber.Ctx) error {
	user := new(structures.Seeker)

	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": err})
	}

	success, message := services.AddSeekerToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(400).JSON(fiber.Map{"message": message})
}

func GetSeeker(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func GetSeekers() (bool, string) {
	success := true
	return success, "users list here"
}

func UpdateSeeker(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func DeleteSeeker(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}
