package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-gin-todo/api/controller"
	"github.com/kritmet/go-gin-todo/app"
	"github.com/kritmet/go-gin-todo/repository"
	"github.com/kritmet/go-gin-todo/usecase"
)

func todoRoute(route *gin.RouterGroup, application *app.Application) {
	uc := usecase.NewTodoUsecase(
		repository.NewTodoRepository(),
	)
	ct := controller.NewTodoController(uc)

	route.GET("", ct.GetAll)
	route.POST("", ct.Create)
	route.PUT("/:id", ct.Update)
}
