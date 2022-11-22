package midlleware

import (
	"github.com/gin-gonic/gin"
)

var err error

func RequireAuth(c *gin.Context) {
	//tokenValue, err := c.Cookie("Authorization")
	//signingKey := []byte(os.Getenv("SECRET"))
	//if err != nil {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//}
	//
	//token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return signingKey, nil
	//})
	//
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	// check time exe
	//	if float64(time.Now().Unix()) > claims["exp"].(float64) {
	//		c.AbortWithStatus(http.StatusUnauthorized)
	//	}
	//	var user models.User
	//
	//	postgres.DB.First(&user, claims["sub"])
	//	if user.UserID == 0 {
	//		c.AbortWithStatus(http.StatusUnauthorized)
	//	}
	//
	//	c.Set("user", user)
	//	c.Next()
	//
	//} else {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//}

}
