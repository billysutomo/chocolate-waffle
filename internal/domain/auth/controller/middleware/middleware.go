package middleware

//AuthMiddleware AuthMiddleware
type AuthMiddleware struct {
}

// InitMiddleware initialize the middleware
func InitMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
