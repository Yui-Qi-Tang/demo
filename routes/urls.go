package routes

import (
	"github.com/gin-gonic/gin"
	"iprotector.github.com/routes/product"
	"iprotector.github.com/routes/user"
	"iprotector.github.com/store"
)

// v1 routes
func V1(store store.Storer) func(router *gin.Engine) {
	return func(router *gin.Engine) {
		v1 := router.Group("/v1")

		v1.Use(
			SayHi,
			// middlewares here...
		)
		user.BindRoute(v1.Group("/user"), store)

		v1WithoutSayHi := router.Group("/v1")
		product.BindRoute(v1WithoutSayHi.Group("/product"), store)
		// yyy.BindRoute
		// ...
		// ...
	}
}
