package models

import "time"

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:20"`
	LastName  string `gorm:"size:20"`
	UserName  string `gorm:"uniqueIndex;size:16"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Contact   string `gorm:"size:13"`
	Contact2  string `gorm:"size:13"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Owner struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:20"`
	LastName  string `gorm:"size:20"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Contact   string `gorm:"size:13"`
	Address   string `gorm:"size:40"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Seeker struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:20"`
	LastName  string `gorm:"size:20"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Contact   string `gorm:"size:13"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tenant struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:20"`
	LastName  string `gorm:"size:20"`
	Email     string `gorm:"uniqueIndex;size:100"`
	Contact   string `gorm:"size:13"`
	HouseRef  string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
