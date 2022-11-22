package service

import (
	"fmt"
	"main.go/dto"
	"main.go/models"
	"main.go/repository"
)

type AuthService interface {
	CreateUser(user dto.UserCreate) models.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) CreateUser(user dto.UserCreate) models.User {
	fmt.Printf("user: %v", user)
	userToCreate := models.User{}
	fmt.Printf("userToC11111111111111111reate", userToCreate)

	fmt.Printf("Creating user with email %s", user.Email)
	fmt.Printf("userToCreate", userToCreate)

	res := service.userRepository.InsertUser(userToCreate)
	fmt.Printf("resssssssssssss:", res)

	return res
}
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
