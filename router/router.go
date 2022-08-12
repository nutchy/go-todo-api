package router

import "github.com/gin-gonic/gin"

func NewGinRouter(routes Routes) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	for _, route := range routes {
		v1.GET(route.Path, route.HandleFunc)
	}

	return router
}
