package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main.go/models"
)

type UserRepository interface {
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	AllUser() []models.User
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	DeleteUser(b models.User)
	FindUserByID(userID uint64) models.User
	VerifyUser(email string, password string) interface{}
	InfoUsers() []models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	user.Role = "user"
	db.connection.Create(&user)
	return user
}

func (db *userConnection) VerifyUser(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		db.connection.Find(&tempUser, user.Id)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}
func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) DeleteUser(b models.User) {
	db.connection.Delete(&b)
}

func (db *userConnection) FindUserByID(userID uint64) models.User {
	var user models.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) AllUser() []models.User {
	var users []models.User
	db.connection.Find(&users)
	return users
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}
func (db *userConnection) InfoUsers() []models.User {

	var users []models.User
	db.connection.Find(&users)
	return users
}
