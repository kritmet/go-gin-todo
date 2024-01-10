package usecase

import (
	"errors"
	"sort"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/kritmet/go-gin-todo/domain"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type todoUsecase struct {
	repository domain.TodoRepository
}

// NewTodoUsecase is a function for create todo usecase
func NewTodoUsecase(repository domain.TodoRepository) domain.TodoUsecase {
	return &todoUsecase{
		repository: repository,
	}
}

// Create is a function for create todo
func (uc *todoUsecase) Create(request *domain.CreateTodoRequest) error {
	entities, err := uc.repository.ReadTodoJSON()
	if err != nil {
		logrus.Errorf("read todo json error: %s", err)
		return err
	}

	entity := &domain.Todo{}
	_ = copier.Copy(entity, request)
	entities = append(entities, entity)

	err = uc.repository.WriteTodoJSON(entities)
	if err != nil {
		logrus.Errorf("write todo json error: %s", err)
		return err
	}

	return nil
}

// GetAll is a function for get all todo
func (uc *todoUsecase) GetAll(request *domain.GetAllRequest) ([]*domain.Todo, error) {
	entities, err := uc.repository.ReadTodoJSON()
	if err != nil {
		logrus.Errorf("read todo json error: %s", err)
		return nil, err
	}

	if request.Title != "" {
		entities = lo.Filter(entities, func(entity *domain.Todo, _ int) bool {
			return strings.Contains(strings.ToLower(entity.Title), strings.ToLower(request.Title))
		})
	}

	if request.Status != "" {
		entities = lo.Filter(entities, func(entity *domain.Todo, _ int) bool {
			return entity.Status == request.Status
		})
	}

	switch o := request.OrderBy; {
	case o.IsTitle():
		sort.Slice(entities, func(i, j int) bool {
			if request.Reverse {
				return entities[i].Title > entities[j].Title
			}
			return entities[i].Title < entities[j].Title
		})

	case o.IsDate():
		sort.Slice(entities, func(i, j int) bool {
			if request.Reverse {
				return entities[i].Date.After(entities[j].Date)
			}
			return entities[i].Date.Before(entities[j].Date)
		})

	case o.IsStatus():
		sort.Slice(entities, func(i, j int) bool {
			if request.Reverse {
				return entities[i].Status > entities[j].Status
			}
			return entities[i].Status < entities[j].Status
		})
	}

	return entities, nil
}

// GetOne is a function for get one todo by id
func (uc *todoUsecase) GetOne(id string) (*domain.Todo, error) {
	entities, err := uc.repository.ReadTodoJSON()
	if err != nil {
		logrus.Errorf("read todo json error: %s", err)
		return nil, err
	}

	entity, found := lo.Find(entities, func(entity *domain.Todo) bool {
		return entity.ID == id
	})
	if !found {
		logrus.Errorf("find todo not found")
		return nil, errors.New("data not found")
	}

	return entity, nil
}

// Update is a function for update todo by id
func (uc *todoUsecase) Update(request *domain.UpdateTodoRequest) error {
	entities, err := uc.repository.ReadTodoJSON()
	if err != nil {
		logrus.Errorf("read todo json error: %s", err)
		return err
	}

	entity, _ := lo.Find(entities, func(entity *domain.Todo) bool {
		return entity.ID == request.ID
	})

	_ = copier.Copy(entity, request)
	err = uc.repository.WriteTodoJSON(entities)
	if err != nil {
		logrus.Errorf("write todo json error: %s", err)
		return err
	}

	return nil
}
