package midlleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"main.go/pkg/helper"
	"main.go/service"
	"net/http"
	"os"
	"strings"
	"time"
)

var err error

func RequireAuth(c *gin.Context) {
	tokenValue, err := c.Cookie("Authorization")
	signingKey := []byte(os.Getenv("SECRET"))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check time exe
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request please Login", "error", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		jwtString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwtService.ValidateToken(jwtString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[id]: ", claims["id"])
			ctx.Next()
		} else {
			respronse := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, respronse)
		}
	}
}
