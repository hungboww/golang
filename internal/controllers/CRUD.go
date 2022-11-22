package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main.go/models"
	"net/http"
)

func GetInfoUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "err",
		})
		return
	}
	uid := user.(models.User).Id
	fmt.Printf("i' i` iiii", uid)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    user,
		"userID": uid,
	})
}
func ListUser(c *gin.Context) {
	//_, exists := c.Get("user")
	//if !exists {
	//	log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": "helpers.MiddlewareErr",
	//	})
	//	return
	//}
	//var accountUser []models.User
	//postgres.DB.Find(&accountUser)
	//c.JSON(http.StatusOK, &accountUser)
}
func DetailUser(c *gin.Context) {}

//	_, exists := c.Get("user")
//	id := c.Params.ByName("id")
//	if !exists {
//		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": "helpers.MiddlewareErr",
//		})
//		return
//	}
//	var accountUser []models.User
//	postgres.DB.First(&accountUser, id)
//	c.JSON(http.StatusOK, &accountUser)
//}
//func UpdateInfo(c *gin.Context) {
//	users, exists := c.Get("user")
//
//	if !exists {
//		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": "err",
//		})
//		return
//	}
//	uid := users.(models.User).UserID
//	var body struct {
//		FirstName string `json:"first_name"`
//		Email     string `json:"email"`
//		Password  string `json:"password"`
//		Lastname  string `json:"last_name"`
//	}
//	fmt.Println("body 1111111", body)
//	if c.Bind(&body) != nil {
//		return
//	}
//	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
//	if err != nil {
//		return
//	}
//	user := models.User{
//		Email:     body.Email,
//		Password:  string(hash),
//		FirstName: body.FirstName,
//		LastName:  body.Lastname,
//		Role:      "user",
//		UpdatedAt: time.Now(),
//	}
//	postgres.DB.Model(&user).Where("user_id = ?", uid).Updates(user)
//
//	c.BindJSON(&user)
//	postgres.DB.Save(&user)
//	c.JSON(200, user)
//}
//func DeletePerson(c *gin.Context) {
//	id, err := strconv.Atoi(c.Params.ByName("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": "err",
//		})
//		return
//	}
//	var user models.User
//	if err := postgres.DB.Where("user_id = ?", id).Delete(&user).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": "err",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"data": user,
//	})
//}
