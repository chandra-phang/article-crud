package main

import (
	"article-crud/app"
	_ "article-crud/docs"
)

// General API info for swaggo/swag library - see: https://github.com/swaggo/swag#general-api-info
//
// @title         Articles CRUD
// @version       0.0.1
// @description   To create, retrieve, update and delete article
// @contact.name  Chandra Phang
// @contact.email chandraphang.idn@gmail.com

func main() {
	application := app.NewApplication()
	application.InitApplication()
}
