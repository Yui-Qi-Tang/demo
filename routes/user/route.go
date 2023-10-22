package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ent "iprotector.github.com/entities/user"
)

// the service that is expected by route
type UserService interface {
	Get(id int) ent.User
}

// route user only cares about the service impl
func New(srv UserService) *UserRoute {
	return &UserRoute{
		srv: srv,
	}
}

type UserRoute struct {
	srv UserService
}

func (r *UserRoute) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r.srv.Get(1))
}
