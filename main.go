package main

import (
	"github.com/kritmet/go-gin-todo/api/route"
	"github.com/kritmet/go-gin-todo/app"
	"github.com/sirupsen/logrus"
)

func main() {
	application, err := app.New()
	if err != nil {
		logrus.Panic(err)
	}

	route.Setup(application)
}
