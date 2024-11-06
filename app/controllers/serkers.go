package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateSeeker(context *fiber.Ctx) (bool, string) {
	user := new(structures.Seeker)

	if err := context.BodyParser(user); err != nil {
		return false, "invalid fields supplied"
	}

	var success = services.AddSeekerToDb(db, user)
	if !success {
		return !success, "failed to create user"
	}
	return success, "created"
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
