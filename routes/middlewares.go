package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SayHi(ctx *gin.Context) {
	fmt.Println("hi!")
	ctx.Next()
}
