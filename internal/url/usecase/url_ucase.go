package usecase

import (
	"github.com/billysutomo/chocolate-waffle/internal/domain"
)

type urlService struct {
	urlRepo domain.URLRepository
}

//NewURLService user service
func NewURLService(a domain.URLRepository) domain.URLService {
	return &urlService{
		urlRepo: a,
	}
}
