package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddSeekerToDb(db *gorm.DB, user *structures.Seeker) (bool, string) {
	suceess := true

	new_user := new(models.Seeker)
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName
	new_user.Email = user.Email
	new_user.Contact = user.Contact
	new_user.Password = user.Password

	db.AutoMigrate(&new_user)

	result := db.Create(&new_user)

	if result.Error != nil {
		return !suceess, result.Error.Error()
	}
	return suceess, "created"
}

func GetSeekerByEmail(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Seeker) {
	var seeker models.Seeker
	
	var email struct{
		Email string
	}

	if err := ctx.BodyParser(&email); err != nil{
		return false, seeker
	}

	err := db.First(&seeker, "email=?", email.Email).Error
	if err != nil {
		return false, seeker
	}
	return true, seeker
}

func UpdateSeekerById(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Seeker) {
	var seeker models.Seeker

	if id, err := ctx.ParamsInt("id"); err == nil {
		if err := db.First(&seeker, id).Error; err != nil {
			return false, seeker
		}
	} else {
		return false, seeker
	}

	var fields map[string]interface{}
	if err := ctx.BodyParser(&fields); err != nil {
		return false, seeker
	}

	if err := db.Model(&seeker).Updates(fields).Error; err != nil {
		return false, seeker
	}
	return true, seeker
}

func DeleteSeekerById(db *gorm.DB, ctx *fiber.Ctx) (bool, string) {
	var admin models.Seeker

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return false, "could not parse id"
	}

	err = db.First(&admin, id).Error
	if err != nil {
		return false, "user not found"
	}

	db.Delete(&models.Seeker{}, id)
	return true, "deleted"
}
