package application

import "article-crud/handlers"

func InitServices(h handlers.Handler) {
	InitArticleService(h)
}
