package service

import (
	"errors"
	"go-todo/model"
	"go-todo/repository"
	"go-todo/utils"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (s *UserService) RegisterUser(user *model.User) error {
	hashedPassword, err := utils.hashedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.UserRepository.CreateUser(user)
	return err
}

func (s *UserService) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := s.UserRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("Invalid Password")
	}
	return user, nil
}
