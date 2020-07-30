package domain

import "context"

// URL url
type URL struct {
	Slug string `json:"slug" validate:"required,slugcheck"`
	URL  string `json:"url" validate:"required"`
}

// URLService service
type URLService struct {
}

// URLRepository repository
type URLRepository interface {
	Create(ctx context.Context)
}
