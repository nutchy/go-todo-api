package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Path       string
	HandleFunc gin.HandlerFunc
}

type Routes []Route

func NewRoutes() Routes {
	var routes = []Route{
		{
			Path: "/hello",
			HandleFunc: func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"firstname": "chayanon",
				})
			},
		},
	}
	return routes
}
