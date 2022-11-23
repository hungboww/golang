package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/dto"
	"main.go/models"
	"main.go/pkg/helper"
	"main.go/service"
	"net/http"
)

type AuThenController interface {
	RegisterAccount(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userController struct {
	userService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuThenController {
	return &userController{
		userService: authService,
		jwtService:  jwtService,
	}
}

func (c *userController) RegisterAccount(ctx *gin.Context) {
	fmt.Println("RegisterAccount", *c)
	var registerBody dto.UserCreate
	err := ctx.ShouldBindJSON(&registerBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("registerBody", registerBody.Email)
	if !c.userService.IsDuplicateEmail(registerBody.Email) {
		fmt.Errorf("Email %s already exists", registerBody.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is duplicate",
		})
		return
	} else {
		createUser := c.userService.CreateUser(registerBody)
		fmt.Println("createUser::::::::;;", createUser)
		res := helper.BuildResponse(true, "Ok", nil, createUser)
		ctx.JSON(http.StatusOK, res)
	}

}
func (c *userController) Login(ctx *gin.Context) {
	var loginOTD dto.Login
	err := ctx.ShouldBind(&loginOTD)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authJWT := c.userService.VerifyTokenUser(loginOTD.Email, loginOTD.Password)
	if v, ok := authJWT.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(string(rune(v.Id)))
		ctx.SetCookie("Authorization", generatedToken, 3600*24*30, "", "", false, true)

		ctx.SetSameSite(http.SameSiteLaxMode)

		//res := helper.BuildResponse(true, "Ok", nil, res)
		res := helper.BuildResponse(true, generatedToken, nil, v)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := helper.BuildErrorResponse("Please check again your credential", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
}

//func RegisterUser(c *gin.Context) {
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
//	fmt.Println("body 2222222", body)
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
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//
//	checkEmail := DB.First(&user, "email = ?", body.Email)
//	if checkEmail == nil {
//		return
//	}
//	result := postgres.DB.Create(&user)
//	if result.Error != nil {
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"status":  http.StatusOK,
//		"message": "success",
//	})
//}
//
//func LoginAccount(c *gin.Context) {
//	var body struct {
//		Email    string `json:"email"`
//		Password string `json:"password"`
//	}
//	if c.Bind(&body) != nil {
//		return
//	}
//
//	var user models.User
//	postgres.DB.First(&user, "Email=?", body.Email)
//	if user.UserID == 0 {
//		return
//	}
//	fmt.Println("user", user.UserID)
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
//	fmt.Println("err", err)
//	if err != nil {
//		return
//	}
//	// Sign and get the complete encoded token as a string using the secret
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"sub": user.UserID,
//		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
//	})
//	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
//	if err != nil {
//		return
//	}
//	c.SetSameSite(http.SameSiteLaxMode)
//	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
//	c.Header("Bearer", tokenString)
//	c.JSON(http.StatusOK, gin.H{
//		"status":  http.StatusOK,
//		"token":   tokenString,
//		"message": "success",
//	})
//}
