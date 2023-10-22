package user

import (
	"github.com/gin-gonic/gin"
	repoUser "iprotector.github.com/repo/user"
	srvUser "iprotector.github.com/service/user"
	"iprotector.github.com/store"
)

func BindRoute(rg *gin.RouterGroup, store store.Storer) error {
	user := New(srvUser.New(repoUser.New(store)))
	rg.GET("", user.Get)
	return nil
}
