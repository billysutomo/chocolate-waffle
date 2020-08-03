package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware AuthMiddleware
type AuthMiddleware struct {
}

// Claims Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func checkToken(myToken string) bool {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	tkn, err := jwt.ParseWithClaims(myToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err == nil && tkn.Valid {
		return true
	}
	return false
}

//AuthRouteMiddleware AuthRouteMiddleware
func (m *AuthMiddleware) AuthRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("X-Auth-Token")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		} else if checkToken(token) {
			c.Next()
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
