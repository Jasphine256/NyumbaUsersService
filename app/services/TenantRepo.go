package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddTenantToDb(db *gorm.DB, user *structures.Tenant) (bool, string) {
	suceess := true

	new_user := new(models.Tenant)
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName
	new_user.Email = user.Email
	new_user.Contact = user.Contact
	new_user.HouseRef = user.HouseRef
	new_user.Password = user.Password

	db.AutoMigrate(&new_user)

	result := db.Create(&new_user)

	if result.Error != nil {
		return !suceess, result.Error.Error()
	}
	return suceess, "created"
}

func GetTenantByEmail(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Tenant) {
	var tenant models.Tenant

	var email struct {
		Email string
	}

	if err := ctx.BodyParser(&email); err != nil {
		return false, tenant
	}

	err := db.First(&tenant, "email=?", email.Email).Error
	if err != nil {
		return false, tenant
	}
	return true, tenant
}

func UpdateTenantById(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Tenant) {
	var tenant models.Tenant

	if id, err := ctx.ParamsInt("id"); err == nil {
		if err := db.First(&tenant, id).Error; err != nil {
			return false, tenant
		}
	} else {
		return false, tenant
	}

	var fields map[string]interface{}
	if err := ctx.BodyParser(&fields); err != nil {
		return false, tenant
	}

	if err := db.Model(&tenant).Updates(fields).Error; err != nil {
		return false, tenant
	}
	return true, tenant
}

func DeleteTenantById(db *gorm.DB, ctx *fiber.Ctx) (bool, string) {
	var admin models.Tenant

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return false, "could not parse id"
	}

	err = db.First(&admin, id).Error
	if err != nil {
		return false, "user not found"
	}

	db.Delete(&models.Tenant{}, id)
	return true, "deleted"
}
