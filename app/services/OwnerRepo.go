package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddOwnerToDb(db *gorm.DB, user *structures.Owner) (bool, string) {
	suceess := true

	new_user := new(models.Owner)
	new_user.FirstName = user.FirstName
	new_user.LastName = user.LastName
	new_user.Email = user.Email
	new_user.Contact = user.Contact
	new_user.Address = user.Address
	new_user.Password = user.Password

	db.AutoMigrate(&new_user)

	result := db.Create(&new_user)

	if result.Error != nil {
		return !suceess, result.Error.Error()
	}
	return suceess, "created"
}

func GetOwnerByEmail(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Owner) {
	var owner models.Owner
	
	var email struct{
		Email string
	}

	if err := ctx.BodyParser(&email); err != nil{
		return false, owner
	}

	err := db.First(&owner, "email=?", email.Email).Error
	if err != nil {
		return false, owner
	}
	return true, owner
}

func UpdateOwnerById(db *gorm.DB, ctx *fiber.Ctx) (bool, models.Owner) {
	var owner models.Owner

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return false, owner
	}

	err = db.First(&owner, id).Error
	if err != nil {
		return false, owner
	}

	var fields map[string]interface{}
	if err = ctx.BodyParser(&fields); err != nil{
		fmt.Println(err)
		return false, owner
	}
	
	err = db.Model(&owner).Updates(fields).Error
	if err != nil {
		fmt.Println(err)
		return false, owner
	}

	return true, owner
}

func DeleteOwnerById(db *gorm.DB, ctx *fiber.Ctx) (bool, string) {
	var admin models.Owner

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return false, "could not parse id"
	}

	err = db.First(&admin, id).Error
	if err != nil {
		return false, "user not found"
	}

	db.Delete(&models.Owner{}, id)
	return true, "deleted"
}
