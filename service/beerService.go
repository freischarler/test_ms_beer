package service

import (
	"github.com/freischarler/hexpattern/domain"
)

type BeerService interface {
	GetAllBeer() ([]domain.Beer, int)
	PostOneBeer(domain.Beer) int
	GetOneByIdBeer(string) (domain.Beer, int)
	GetBoxBeer(string, string, int) (float64, int)
}

type DefaultBeerService struct {
	repo domain.BeerRepository
}

func (s DefaultBeerService) GetAllBeer() ([]domain.Beer, int) {
	return s.repo.GetAll()
}

func (s DefaultBeerService) PostOneBeer(b domain.Beer) int {
	return s.repo.PostOne(b)
}

func (s DefaultBeerService) GetOneByIdBeer(id string) (domain.Beer, int) {
	return s.repo.GetOneByID(id)
}

func NewBeerService(repository domain.BeerRepository) DefaultBeerService {
	return DefaultBeerService{repository}
}

func (s DefaultBeerService) GetBoxBeer(id string, currency string, count int) (float64, int) {
	return s.repo.GetBoxPrice(id, currency, count)
}
