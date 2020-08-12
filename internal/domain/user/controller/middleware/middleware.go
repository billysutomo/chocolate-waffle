package middleware

//UserMiddleware UserMiddleware
type UserMiddleware struct {
}

// InitMiddleware initialize the middleware
func InitMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}
