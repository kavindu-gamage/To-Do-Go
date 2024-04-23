package repository

import (
	"example.com/hello/Documents/SE-Projects/go-todo/model"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user *model.User) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.User{}).Where("email=?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.DB.Model(&model.User{}).Where("username=?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}
