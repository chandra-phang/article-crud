package app

import (
	"article-crud/api"
	"article-crud/application"
	"article-crud/log"
	"context"
)

type Application struct {
}

// Returns a new instance of the application
func NewApplication() Application {
	return Application{}
}

func (a Application) InitApplication() {
	startupCtx := context.Background()
	log.Infof(startupCtx, "Application is starting up")

	application.InitServices()
	api.InitRoutes()
}
