package middleware

import "github.com/gin-gonic/gin"

//GoMiddleware GoMiddleware
type GoMiddleware struct {
}

// // CORS will handle the CORS middleware
// func (m *GoMiddleware) CORS(next gin.HandlerFunc) gin.HandlerFunc {
// 	return func(c gin.Context) error {
// 		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
// 		return next(c)
// 	}
// }

//CORSMiddleware CORSMiddleware
func (m *GoMiddleware) CORSMiddleware() gin.HandlerFunc {
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

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
