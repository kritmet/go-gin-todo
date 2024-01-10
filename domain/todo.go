package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Status is todo status
type Status string

const (
	// StatusInProgress is todo status in progress
	StatusInProgress Status = "IN_PROGRESS"
	// StatusCompleted is todo status completed
	StatusCompleted Status = "COMPLETED"
)

// Todo todo struct
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Image       string    `json:"image"`
	Status      Status    `json:"status"`
}

// CreateTodoRequest create todo request
type CreateTodoRequest struct {
	ID string `json:"id" binding:"required,uuid"`
	TodoRequest
}

// UpdateTodoRequest update todo request
type UpdateTodoRequest struct {
	ID string `json:"-"`
	TodoRequest
}

// TodoRequest todo request
type TodoRequest struct {
	Title       string    `json:"title" binding:"required,max=100"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" time_format:"2006-01-02T15:04:05Z07:00"`
	Image       string    `json:"image" binding:"base64"`
	Status      Status    `json:"status" binding:"required,oneof=IN_PROGRESS COMPLETED"`
}

// GetAllRequest get all request
type GetAllRequest struct {
	Title   string  `form:"title"`
	Status  Status  `form:"status"`
	OrderBy OrderBy `form:"order_by"`
	Reverse bool    `form:"reverse"`
}

// OrderBy is todo order by
type OrderBy string

const (
	// OrderByTitle is todo order by title
	OrderByTitle OrderBy = "title"
	// OrderByDate is todo order by date
	OrderByDate OrderBy = "date"
	// OrderByStatus is todo order by status
	OrderByStatus OrderBy = "status"
)

// IsTitle is a function for check order by is title
func (o OrderBy) IsTitle() bool {
	return o == OrderByTitle
}

// IsDate is a function for check order by is date
func (o OrderBy) IsDate() bool {
	return o == OrderByDate
}

// IsStatus is a function for check order by is status
func (o OrderBy) IsStatus() bool {
	return o == OrderByStatus
}

// TodoController todo controller interface
type TodoController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
}

// TodoUsecase todo usecase interface
type TodoUsecase interface {
	Create(request *CreateTodoRequest) error
	GetAll(request *GetAllRequest) ([]*Todo, error)
	GetOne(id string) (*Todo, error)
	Update(request *UpdateTodoRequest) error
}

// TodoRepository todo repository interface
type TodoRepository interface {
	WriteTodoJSON(entities []*Todo) error
	ReadTodoJSON() ([]*Todo, error)
}
