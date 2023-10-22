package product

import (
	"fmt"
	"time"

	entity "iprotector.github.com/entities/product"
	"iprotector.github.com/store"
)

func New(store store.Storer) *ProductRepository {
	return &ProductRepository{
		store: store,
	}
}

type ProductRepository struct {
	store store.Storer
}

func (pr *ProductRepository) Get(id int) entity.Product {
	msg := pr.store.Get(1)
	fmt.Println(msg)
	return entity.Product{
		ID:        1,
		Name:      "Play Station 5 - Slim",
		Price:     100,
		CreatedAt: time.Now(),
	}
}
