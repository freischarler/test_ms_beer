package service

import (
	"github.com/freischarler/hexpattern/domain"
)

type BeerService interface {
	GetAllBeer() ([]domain.Beer, error)
	PostOneBeer(domain.Beer) error
	GetOneByIdBeer(string) (domain.Beer, error)
	GetBoxBeer(string, string, int) (float64, error)
}

type DefaultBeerService struct {
	repo domain.BeerRepository
}

func (s DefaultBeerService) GetAllBeer() ([]domain.Beer, error) {
	return s.repo.GetAll()
}

func (s DefaultBeerService) PostOneBeer(b domain.Beer) error {
	return s.repo.PostOne(b)
}

func (s DefaultBeerService) GetOneByIdBeer(id string) (domain.Beer, error) {
	return s.repo.GetOneByID(id)
}

func NewBeerService(repository domain.BeerRepository) DefaultBeerService {
	return DefaultBeerService{repository}
}

func (s DefaultBeerService) GetBoxBeer(id string, currency string, count int) (float64, error) {
	return s.repo.GetBoxPrice(id, currency, count)
}
