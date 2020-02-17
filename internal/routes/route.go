package routes

import (
	"github.com/Api-Marvel/internal/controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("index", controllers.GetIndex)
		v1.GET("getcomics", controllers.GetComics)
		v1.POST("createcomic", controllers.CreateComic)
	}
	return router
}
