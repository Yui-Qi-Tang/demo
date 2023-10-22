// Package restapi sets up the api server and start it
// The implementation of the Handler of http.Server is the go-gin api framework
package restapi

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type BindRouter func(router *gin.Engine)

func New(addr string, routes ...BindRouter) (*http.Server, error) {
	router := gin.Default()

	// apply the routes
	for _, r := range routes {
		r(router)
	}

	s := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	return s, nil
}

func Start(ctx context.Context, s *http.Server) error {
	if s == nil {
		return errors.New("server is nil")
	}
	log.Println("start server at", s.Addr)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	log.Println("server is shuting down")
	return s.Shutdown(ctx)
}
