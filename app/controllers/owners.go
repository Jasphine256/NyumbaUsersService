package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateOwner(context *fiber.Ctx) (bool, string) {
	user := new(structures.Owner)

	if err := context.BodyParser(user); err != nil {
		return false, "invalid fields supplied"
	}

	var success = services.AddOwnerToDb(db, user)
	if !success {
		return !success, "failed to create user"
	}
	return success, "created"
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
