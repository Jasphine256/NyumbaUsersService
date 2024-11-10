package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddAdminToDb(db *gorm.DB, user *structures.Admin) (bool, string) {
	suceess := true

	new_user := new(models.Admin)
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName
	new_user.UserName = user.UserName
	new_user.Email = user.Email
	new_user.Contact = user.Contact
	new_user.Contact2 = user.Contact2
	new_user.Password = user.Password

	db.AutoMigrate(&new_user)

	result := db.Create(&new_user)

	if result.Error != nil {
		return !suceess, result.Error.Error()
	}
	return suceess, "created"
}

func GetAdminByEmail(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Admin) {
	var admin models.Admin
	
	var email struct{
		Email string
	}

	if err := ctx.BodyParser(&email); err != nil{
		return false, admin
	}

	err := db.First(&admin, "email=?", email.Email).Error
	if err != nil {
		return false, admin
	}
	return true, admin
}

func UpdateAdminById(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Admin) {
	var admin models.Admin

	if id, err := ctx.ParamsInt("id"); err == nil {
		if err := db.First(&admin, id).Error; err != nil {
			return false, admin
		}
	} else {
		return false, admin
	}

	var fields map[string]interface{}
	if err := ctx.BodyParser(&fields); err != nil {
		return false, admin
	}

	if err := db.Model(&admin).Updates(fields).Error; err != nil {
		return false, admin
	}
	return true, admin
}

func DeleteAdminById(db *gorm.DB, ctx *fiber.Ctx) (bool, string) {
	var admin models.Admin

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return false, "could not parse id"
	}

	err = db.First(&admin, id).Error
	if err != nil {
		return false, "user not found"
	}

	db.Delete(&models.Admin{}, id)
	return true, "deleted"
}
