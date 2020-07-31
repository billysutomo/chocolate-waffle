package domain

// URL url
type URL struct {
	Slug string `json:"slug" validate:"required,slugcheck"`
	URL  string `json:"url" validate:"required"`
}

// URLUsecase service
type URLUsecase interface {
}

// URLRepository repository
type URLRepository interface {
	// Create(ctx context.Context)
}
