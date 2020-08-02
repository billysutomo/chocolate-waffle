package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware AuthMiddleware
type AuthMiddleware struct {
}

// // CORS will handle the CORS middleware
// func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c gin.Context) error {
// 		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
// 		return next(c)
// 	}
// }

// Claims Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func checkToken(myToken string) bool {

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(myToken, claims, func(token *jwt.Token) (interface{}, error) {
		aa := jwt.New(jwt.SigningMethodHS256)
		aa.SignedString([]byte("secret"))
		return aa, nil
	})

	// token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
	// 	key, err := ioutil.ReadFile("your-private-key.pem")
	// 	if err != nil {
	// 		return nil, errors.New("private key could not be loaded")
	// 	}
	// 	return key, nil
	// })

	if err == nil && tkn.Valid {
		return true
	}

	return false
}

// TODO: fixing Headers were already written. Wanted to override status code 200 with 401

//AuthRouteMiddleware AuthRouteMiddleware
func (m *AuthMiddleware) AuthRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }

		token := c.Request.Header.Get("X-Auth-Token")

		if token == "" {
			c.AbortWithStatus(401)
		} else if checkToken(token) {
			c.Next()
		} else {
			c.AbortWithStatus(401)
		}
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
