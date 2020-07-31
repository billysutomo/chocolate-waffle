package model

// URL url
type URL struct {
	Slug string `json:"slug" validate:"required"`
	URL  string `json:"url" validate:"required"`
}

// ResponseError response error
type ResponseError struct {
	Message string `json:"message"`
}
