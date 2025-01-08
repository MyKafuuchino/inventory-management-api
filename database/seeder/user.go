package seeder

import (
	"golang.org/x/crypto/bcrypt"
	"inventory-management/database"
	"inventory-management/entity"
	"log"
)

func SeedUser() {
	var existingAdmin entity.User
	err := database.DB.Where("username = ?", "admin").First(&existingAdmin).Error
	if err == nil {
		log.Printf("Admin user already exists with username: %s", existingAdmin.Username)
		return
	}

	admin := entity.User{
		Username: "admin",
		FullName: "Administrator",
		Password: "admin",
		Role:     "admin",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password for user %s: %v", admin.Username, err)
	}
	admin.Password = string(hashedPassword)

	err = database.DB.Create(&admin).Error
	if err != nil {
		log.Printf("Failed to create user %s: %v", admin.Username, err)
	} else {
		log.Printf("User %s created successfully", admin.Username)
	}
}
