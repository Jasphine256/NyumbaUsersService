package services

import (
	"NyumbaUsersService/app/models"
	"NyumbaUsersService/app/structures"
	"log"

	"gorm.io/gorm"
)

// ################# ADD USERS TO DB HANDLERS #####################

func AddAdminToDb(db *gorm.DB, user *structures.Admin) bool {
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
		log.Fatalf("Failed to create user: %v", result.Error)
		return !suceess
	}
	return suceess
}

func AddOwnerToDb(db *gorm.DB, user *structures.Owner) bool {
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
		return !suceess
	}
	return suceess
}

func AddSeekerToDb(db *gorm.DB, user *structures.Seeker) bool {
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
		log.Fatalf("Failed to create user: %v", result.Error)
		return !suceess
	}
	return suceess
}

func AddTenantToDb(db *gorm.DB, user *structures.Tenant) bool {
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
		log.Fatalf("Failed to create user: %v", result.Error)
		return !suceess
	}
	return suceess
}
