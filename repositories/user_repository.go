package repositories

import (
	"golang-fiber-postgres-template/config"
	"golang-fiber-postgres-template/models"
)

// CreateUser inserts a new user in the database.
func CreateUser(u *models.User) error {
	return config.DB.Create(u).Error
}

// GetUserByUsername returns a user by username.
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
