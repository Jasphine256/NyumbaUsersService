package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"
	"log"

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
		log.Fatalf("Failed to create user: %v", result.Error)
		return !suceess, "failed"
	}
	return suceess, "created"
}
