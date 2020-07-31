package usecase

import (
	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type urlUsecase struct {
	urlRepo domain.URLRepository
}

//NewURLUsecase user service
func NewURLUsecase(a domain.URLRepository) domain.URLUsecase {
	return &urlUsecase{
		urlRepo: a,
	}
}
