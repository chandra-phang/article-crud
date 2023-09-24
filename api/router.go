package api

import (
	v1 "article-crud/api/controller/v1"
	"article-crud/apiconst"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type router struct {
	// lightweight and fast HTTP router for Go web applications
	server *chi.Mux
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	return r
}

func InitRoutes() {
	r := newRouter()

	articleController := v1.InitArticleController()
	r.Route("/v1", func(r chi.Router) {
		r.Get("/articles", articleController.ListArticles)
		r.Get("/articles/{id}", articleController.GetArticle)
		r.Post("/articles", articleController.CreateArticle)
		r.Put("/articles/{id}", articleController.UpdateArticle)
		r.Delete("/articles{id}", articleController.DeleteArticle)
	})

	// Mount Swagger docs and redirect / to the Swagger UI
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.UIConfig(map[string]string{
			"onComplete": fmt.Sprintf(`() => { document.title = '%s'; }`, apiconst.ServiceDocsTitle),
		}),
	))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":3000", r)
}
