package product

import (
	entity "iprotector.github.com/entities/product"
)

type ProductRepository interface {
	Get(id int) entity.Product
}

func New(repo ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

type ProductService struct {
	repo ProductRepository
}

func (ps *ProductService) Get(id int) entity.Product {
	return ps.repo.Get(1)
}
