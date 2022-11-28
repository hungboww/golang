package controllers

import (
	"github.com/gin-gonic/gin"
	"main.go/dto"
	"main.go/models"
	"main.go/pkg/helper"
	"main.go/service"
	"net/http"
	"strconv"
)

type UserController interface {
	AllUserControllers(ctx *gin.Context)
	DetailUserControllers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}
type userControllers struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewAccountController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userControllers{
		userService: userService,
		jwtService:  jwtService,
	}
}
func (c *userControllers) AllUserControllers(ctx *gin.Context) {
	var account []models.User = c.userService.AllUser()
	res := helper.BuildResponse(true, "Ok", nil, account)
	ctx.JSON(http.StatusOK, res)
}
func (c *userControllers) DetailUserControllers(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var account models.User = c.userService.GetUserById(id)
	if (account == models.User{}) {
		res := helper.BuildErrorResponse("No account was found", "No account was found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	} else {
		res := helper.BuildResponse(true, "Ok", nil, account)
		ctx.JSON(http.StatusOK, res)
	}

}
func (c *userControllers) DeleteUser(ctx *gin.Context) {
	var account models.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	account = c.userService.GetUserById(id)
	if account.Id == id {
		c.userService.DeleteUser(account)
		res := helper.BuildResponse(true, "Delete success!@", nil, helper.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("No account was found", "No account was found", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

}
func (c *userControllers) UpdateUser(ctx *gin.Context) {
	var userUpdate dto.UserUpdate
	err := ctx.ShouldBind(&userUpdate)

	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.GetHeader("Authorization")
	result := c.userService.UpdateUser(userUpdate)
	response := helper.BuildResponse(true, "OK", nil, result)
	ctx.JSON(http.StatusOK, response)

}
func GetInfoUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "err",
		})
		return
	}
	uid := user.(models.User).Id
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    user,
		"userID": uid,
	})
}

//func ListUser(c *gin.Context) {
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
//}
//func DetailUser(c *gin.Context) {}

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
