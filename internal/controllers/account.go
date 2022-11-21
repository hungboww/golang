package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"main.go/models"
	"main.go/pkg/db/postgres"
	"net/http"
	"os"
	"time"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Lastname  string `json:"last_name"`
	}
	if c.Bind(&body) != nil {
		return
	}
	fmt.Println("body", body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return
	}
	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		FirstName: body.FirstName,
		LastName:  body.Lastname,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	checkEmail := postgres.DB.First(&user, "email = ?", body.Email)
	if checkEmail == nil {
		return
	}
	fmt.Println("11111", body.Lastname)
	result := postgres.DB.Create(&user)
	fmt.Println("result:", result)
	if result.Error != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

func LoginAccount(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {

		return
	}
	var user models.User
	postgres.DB.First(&user, "Email=?", body.Email)
	if user.UserID == 0 {
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return
	}
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.Header("Bearer", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"token":   tokenString,
		"message": "success",
	})
}
