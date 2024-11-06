package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateAdmin(context *fiber.Ctx) (bool, string) {
	user := new(structures.Admin)

	if err := context.BodyParser(user); err != nil {
		return false, "invalid fields supplied"
	}

	var success = services.AddAdminToDb(db, user)
	if !success {
		return !success, "failed to created admin"
	}
	return success, "created"
}

func GetAdmin(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func GetAdmins() (bool, string) {
	success := true
	return success, "users list here"
}

func UpdateAdmin(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}

func DeleteAdmin(context *fiber.Ctx) (bool, string) {
	success := true
	return success, "data"
}
