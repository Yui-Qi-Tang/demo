package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "iprotector.github.com/entities/product"
)

type ProductService interface {
	Get(id int) entity.Product
}

func New(srv ProductService) *ProductRoute {
	return &ProductRoute{
		productSrv: srv,
	}
}

type ProductRoute struct {
	productSrv ProductService
}

func (r *ProductRoute) Get(c *gin.Context) {
	c.JSON(http.StatusOK, r.productSrv.Get(1))
}
