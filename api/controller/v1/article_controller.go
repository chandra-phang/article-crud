package v1

import (
	"article-crud/api/controller"
	"article-crud/application"
	"net/http"
)

type articleController struct {
	svc application.IArticleService
}

// creates a new instance of this controller with reference to ArticleService
func InitArticleController() *articleController {
	//  initializes its "svc" field with a service instance returned by "application.GetArticleService()".
	return &articleController{
		svc: application.GetArticleService(),
	}
}

func (c *articleController) ListArticles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c.svc.ListArticles(ctx)

	controller.WriteSuccess(ctx, w, r, 200, nil)
}

func (c *articleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c.svc.ListArticles(ctx)

	controller.WriteSuccess(ctx, w, r, 200, nil)
}

func (c *articleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c.svc.ListArticles(ctx)

	controller.WriteSuccess(ctx, w, r, 200, nil)
}

func (c *articleController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c.svc.ListArticles(ctx)

	controller.WriteSuccess(ctx, w, r, 200, nil)
}

func (c *articleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	c.svc.ListArticles(ctx)

	controller.WriteSuccess(ctx, w, r, 200, nil)
}
