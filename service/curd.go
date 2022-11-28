package service

import (
	"github.com/mashingan/smapping"
	"log"
	"main.go/dto"
	"main.go/models"
	"main.go/repository"
)

type UserService interface {
	AllUser() []models.User
	GetUserById(id int) models.User
	UpdateUser(user dto.UserUpdate) models.User
	DeleteUser(id models.User)
	Info(id int) models.User
	InfoUsers() []models.User
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) AllUser() []models.User {
	return s.userRepository.AllUser()
}

func (s *userService) GetUserById(id int) models.User {
	return s.userRepository.FindUserByID(uint64(id))
}

func (s *userService) UpdateUser(user dto.UserUpdate) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := s.userRepository.UpdateUser(userToUpdate)
	return res
}

func (s *userService) DeleteUser(id models.User) {
	s.userRepository.DeleteUser(id)
}

func (s *userService) Info(id int) models.User {
	return s.userRepository.FindUserByID(uint64(id))
}

func (s *userService) InfoUsers() []models.User {
	return s.userRepository.InfoUsers()
}
