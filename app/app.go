package app

import (
	"github.com/kritmet/go-gin-todo/app/config"
	"github.com/kritmet/go-gin-todo/docs"
)

// Application application
type Application struct {
	Config *config.Configs
}

// New new application
func New() (*Application, error) {
	err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	app := &Application{
		Config: config.CF,
	}

	docs.SwaggerInfo.Title = config.CF.Swagger.Title
	docs.SwaggerInfo.Description = config.CF.Swagger.Description

	return app, nil
}
