package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "url-shortener/docs" // This line is necessary for go-swagger to find your docs
)

func AddSwaggerRoutes(context *gin.Engine) {
	context.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
