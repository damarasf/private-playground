package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/:shorturl", urlHandler.Get)
	router.POST("/", urlHandler.Create)
	router.POST("/custom", urlHandler.CreateCustom)
	return router
}
