package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritmet/go-gin-todo/domain"
)

// NewTodoController is a function for create todo controller
func NewTodoController(usecase domain.TodoUsecase) domain.TodoController {
	return &todoController{
		usecase: usecase,
	}
}

type todoController struct {
	usecase domain.TodoUsecase
}

// Create Create
// @Tags Todo
// @Summary Create
// @Description create todo data
// @Accept json
// @Produce json
// @Param request body domain.CreateTodoRequest true "input create request"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /api/v1/todo [post]
func (c *todoController) Create(ctx *gin.Context) {
	var request domain.CreateTodoRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewMessage(err))
		return
	}

	err = c.usecase.Create(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewMessage(err))
		return
	}

	ctx.JSON(http.StatusOK, domain.NewMessage())
}

// GetAll GetAll
// @Tags Todo
// @Summary GetAll
// @Description get all todo data
// @Accept json
// @Produce json
// @Param request query domain.GetAllRequest true "query for get all"
// @Success 200 {array} domain.Todo
// @Failure 400 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /api/v1/todo [get]
func (c *todoController) GetAll(ctx *gin.Context) {
	var request domain.GetAllRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewMessage(err))
		return
	}

	entities, err := c.usecase.GetAll(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewMessage(err))
		return
	}

	ctx.JSON(http.StatusOK, entities)
}

// Update Update
// @Tags Todo
// @Summary Update
// @Description update todo data
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param request body domain.UpdateTodoRequest true "input update request"
// @Success 200 {object} domain.Message
// @Failure 400 {object} domain.Message
// @Failure 404 {object} domain.Message
// @Failure 500 {object} domain.Message
// @Router /api/v1/todo/{id} [put]
func (c *todoController) Update(ctx *gin.Context) {
	var request domain.UpdateTodoRequest
	request.ID = ctx.Param("id")
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.NewMessage(err))
		return
	}

	_, err = c.usecase.GetOne(request.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.NewMessage(err))
		return
	}

	err = c.usecase.Update(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewMessage(err))
		return
	}

	ctx.JSON(http.StatusOK, domain.NewMessage())
}
