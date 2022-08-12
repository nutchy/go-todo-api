package main

import (
	"flag"
	"fmt"

	"github.com/nutchy/go-todo-api/configs"
	"github.com/nutchy/go-todo-api/router"
)

func main() {

	port := flag.String("port", "8021", "running port")

	if err := configs.NewConfig(); err != nil {
		panic(err)
	}

	routes := router.NewRoutes()
	router := router.NewGinRouter(routes)

	router.Run(fmt.Sprintf(":%s", *port))
}
