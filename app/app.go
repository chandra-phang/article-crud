package app

import (
	"article-crud/api"
	"article-crud/application"
	"article-crud/db"
	"article-crud/handlers"
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

	database := db.Connect()
	h := handlers.New(database)

	application.InitServices(h)
	api.InitRoutes()

	db.CloseConnection(database)
}
