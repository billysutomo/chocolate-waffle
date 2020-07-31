package usecase

import (
	"context"
	"log"

	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type urlUsecase struct {
	urlRepo domain.URLRepository
}

// NewURLUsecase user service
func NewURLUsecase(a domain.URLRepository) domain.URLUsecase {
	return &urlUsecase{
		urlRepo: a,
	}
}

// GetURL GetURL
func (a *urlUsecase) GetURL(c context.Context, nama string) {
	log.Printf(nama)
	a.urlRepo.GetURL(c, "sutomo")
	return
}
