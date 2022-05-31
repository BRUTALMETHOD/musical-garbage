package routers

import (
	"github.com/brutalmethod/musicalgarbage/controllers"
	"github.com/brutalmethod/musicalgarbage/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	m := middlewares.NewMetrics()

	v1 := router.Group("v1")
	m.Use(v1)
	{
		artGroup := v1.Group("art")
		{
			artGroup.GET("/nodename", gin.WrapF(controllers.Draw))
		}

	}
	return router

}
