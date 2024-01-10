package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-gin-todo/app"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup setup route
func Setup(application *app.Application) {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")
	todo := v1.Group("/todo")
	todoRoute(todo, application)

	err := router.Run(fmt.Sprintf(":%d", application.Config.App.Port))
	if err != nil {
		panic(err)
	}
}
