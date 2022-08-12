package main

import (
	"flag"
	"fmt"

	"github.com/nutchy/go-todo-api/router"
)

func main() {

	port := flag.String("port", "8021", "running port")

	routes := router.NewRoutes()
	router := router.NewGinRouter(routes)

	router.Run(fmt.Sprintf(":%s", *port))
}
