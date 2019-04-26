package listing

import (
	"errors"
)

var ErrNotFound = errors.New("product not found")

type Repository interface {
	Product(int) (Product, error)
	Products() []Product
}

type Service interface {
	Product(int) (Product, error)
	Products() []Product
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Products() []Product {
	return s.r.Products()
}

func (s *service) Product(id int) (Product, error) {
	return s.r.Product(id)
}
