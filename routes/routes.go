package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", RootRouteController)
	server.GET("/health", HealthCheckController)
	server.POST("/shorten", shortenURL)
	server.GET("/:shortURL", redirectToOriginalURL)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
