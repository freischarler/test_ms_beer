package domain

type Beer struct {
	Id       int
	Name     string
	Brewery  string
	Country  string
	Price    float64
	Currency string
}

type BeerBox struct {
	PriceTotal int
}

type BeerRepository interface {
	GetAll() ([]Beer, error)
	PostOne(Beer) error
	GetOneByID(string) (Beer, error)
	GetBoxPrice(string, string, int) (float64, error)
}
