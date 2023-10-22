package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"iprotector.github.com/restapi"
	"iprotector.github.com/routes"
)

type fakeStore string

func (f fakeStore) Get(id any) any {
	return fmt.Sprintf("get data from store '%s' with id %v", f, id)
}

// main handles the 3rd party modules and controls/combines them into restapi server
// 3rd party modules includes:
//   - database/persistent volume...
//   - cache: in-memory or persistent
//   - message broker server
func main() {
	addr := ":8080"
	var fs fakeStore = "fake"
	server, err := restapi.New(addr, routes.V1(fs))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := restapi.Start(ctx, server); err != nil {
		panic(err)
	}
	log.Println("server has been shutdown completed")
}
