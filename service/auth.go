package service

import (
	"fmt"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
	"main.go/dto"
	"main.go/models"
	"main.go/repository"
)

type AuthService interface {
	CreateUser(user dto.UserCreate) models.User
	IsDuplicateEmail(email string) bool
	VerifyTokenUser(email string, password string) interface{}
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
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}
func (service *authService) IsDuplicateEmail(email string) bool {
	fmt.Println("email user :", email)
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) VerifyTokenUser(email string, password string) interface{} {
	res := service.userRepository.VerifyUser(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
