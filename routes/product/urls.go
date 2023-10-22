package product

import (
	"github.com/gin-gonic/gin"
	repoProduct "iprotector.github.com/repo/product"
	srvProduct "iprotector.github.com/service/product"
	"iprotector.github.com/store"
)

func BindRoute(rg *gin.RouterGroup, store store.Storer) error {
	product := New(srvProduct.New(repoProduct.New(store)))
	rg.GET("", product.Get)
	return nil
}
