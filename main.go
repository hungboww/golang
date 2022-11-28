package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/CSV"
	"main.go/internal/controllers"
	"main.go/midlleware"
	"main.go/pkg/db/postgres"
	"main.go/repository"
	"main.go/service"
)

var (
	db                *gorm.DB                     = postgres.Connection()
	csvData           CSV.AccountData              = CSV.NewAccountCSVData(db)
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	authController    controllers.AuThenController = controllers.NewAuthController(authService, jwtService)
	jwtService        service.JWTService           = service.NewJWTService()
	accountService    service.UserService          = service.NewUserService(userRepository)
	accountController controllers.UserController   = controllers.NewAccountController(accountService, jwtService)
)

func main() {
	defer postgres.CloseDatabaseConnection(db)
	r := gin.Default()
	rows2 := CSV.ReadFile("account.csv")
	csvData.InsertDataAccount(rows2)
	authRoutes := r.Group("api/account")
	{
		authRoutes.POST("/signup", authController.RegisterAccount)
		authRoutes.POST("/login", authController.Login)
	}
	userRoutes := r.Group("api/user", midlleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/all", accountController.AllUserControllers)
		userRoutes.GET("/profile/:id", accountController.DetailUserControllers)
		userRoutes.DELETE("/profile/:id", accountController.DeleteUser)
		userRoutes.PATCH("/profile/:id", accountController.UpdateUser)
	}
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
