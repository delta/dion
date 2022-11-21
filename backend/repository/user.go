package repository

import (
	"delta.nitt.edu/dion/models"
	"gorm.io/gorm/clause"
)

func GetUser(email string) (models.User, error) {
	var user models.User
	res := db.First(&user, "email = ?", email)
	return user, res.Error
}

func InsertUser(name string, email string) error {
	user := models.User{Name: name, Email: email}
	return db.Create(&user).Error
}

func UpsertUser(name string, email string) error {
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{Name: name, Email: email})
	return result.Error
}

func DeleteUser(user *models.User) error {
	return db.Delete(user).Error
}
