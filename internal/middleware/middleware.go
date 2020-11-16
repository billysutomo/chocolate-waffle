package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//MainMiddleware GoMiddleware
type MainMiddleware struct {
}

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
		const bearerSchema = "Bearer "
		authHeader := c.Request.Header.Get("Authorization")
		val := strings.HasPrefix(authHeader, bearerSchema)
		if !val {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[len(bearerSchema):]
		claims, ok := checkToken(token)

		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		id := claims["id"].(float64)
		c.Set("id", int(id))
		c.Next()
		return
	}
}

func checkToken(myToken string) (jwt.MapClaims, bool) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	tkn, err := jwt.ParseWithClaims(myToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	data := tkn.Claims.(jwt.MapClaims)

	if err == nil && tkn.Valid {
		return data, true
	}

	return nil, false
}

// InitMiddleware initialize the middleware
func InitMiddleware() *MainMiddleware {
	return &MainMiddleware{}
}
