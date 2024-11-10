package controllers

import (
	"NyumbaUsersService/app/database"
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB = database.GetDB()

func CreateAdmin(context *fiber.Ctx) error {

	user := new(structures.Admin)
	if err := context.BodyParser(user); err != nil {
		return context.Status(400).JSON(fiber.Map{"message": "invalid fields provided"})
	}

	success, message := services.AddAdminToDb(db, user)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": message})
	}
	return context.Status(201).JSON(fiber.Map{"message": message})
}

func GetAdmin(context *fiber.Ctx) error {
	success, data := services.GetAdminByEmail(db, context)
	if !success{
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func UpdateAdmin(context *fiber.Ctx) error {
	success, data := services.UpdateAdminById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}

func DeleteAdmin(context *fiber.Ctx) error {
	success, data := services.DeleteAdminById(db, context)
	if !success {
		return context.Status(400).JSON(fiber.Map{"message": data})
	}
	return context.Status(200).JSON(fiber.Map{"message": data})
}
