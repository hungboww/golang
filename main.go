package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/CSV"
	"main.go/internal/controllers"
	"main.go/pkg/db/postgres"
)

var (
	db *gorm.DB
)

func init() {
	postgres.Connection()
}
func main() {
	r := gin.Default()
	rows2 := CSV.ReadFile("account.csv")
	CSV.InsertDataAccount(rows2)
	userRoutes := r.Group("api/user")
	{
		userRoutes.POST("/signup", controllers.RegisterUser)
		userRoutes.PUT("/login", controllers.LoginAccount)
	}
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
