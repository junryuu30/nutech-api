package repositories

import (
	"nutech/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	// UpdateUser(user models.User, ID int) (models.User, error)
	// DeleteUser(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) FindUsers() ([]models.User, error) {
// 	var users []models.User
// 	// Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
// 	err := r.db.Preload("Product.User").Find(&users).Error // add this code
// 	return users, err
// }

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

// func (r *repository) GetUser(ID int) (models.User, error) {
// 	var user models.User
// 	err := r.db.Preload("Product.User").First(&user, ID).Error

// 	return user, err
// }

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
