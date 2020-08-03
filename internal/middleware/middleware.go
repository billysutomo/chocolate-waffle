package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//MainMiddleware GoMiddleware
type MainMiddleware struct {
}

// // CORS will handle the CORS middleware
// func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c gin.Context) error {
// 		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
// 		return next(c)
// 	}
// }

//CORSMiddleware CORSMiddleware
func (m *MainMiddleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//AuthRouteMiddleware AuthRouteMiddleware
func (m *MainMiddleware) AuthRouteMiddleware() gin.HandlerFunc {
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

// InitMiddleware initialize the middleware
func InitMiddleware() *MainMiddleware {
	return &MainMiddleware{}
}
