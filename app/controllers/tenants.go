package controllers

import (
	"NyumbaUsersService/app/services"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
)

func CreateTenant(context *fiber.Ctx) (bool, string) {
	user := new(structures.Tenant)

	if err := context.BodyParser(user); err != nil {
		return false, "invalid fields supplied"
	}

	var success = services.AddTenantToDb(db, user)
	if !success {
		return !success, "failed to create user"
	}
	return success, "created"
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
