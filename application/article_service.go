package application

import (
	"article-crud/domain/article"
	v1request "article-crud/dto/request/v1"
	"context"
)

type IArticleService interface {
	// svc CRUD methods for domain objects
	ListArticles(ctx context.Context) []article.Article
	GetArticle(ctx context.Context, id string) *article.Article
	CreateArticle(ctx context.Context, dto v1request.CreateArticleDTO) error
	UpdateArticle(ctx context.Context, dto v1request.UpdateArticleDTO, id string) error
	DeleteArticle(ctx context.Context, id string) error
}

type articleSvc struct {
}

var articleSvcSingleton IArticleService

func InitArticleService() {
	articleSvcSingleton = articleSvc{}
}

func GetArticleService() IArticleService {
	return articleSvcSingleton
}

func (svc articleSvc) ListArticles(ctx context.Context) []article.Article {
	return []article.Article{}
}

func (svc articleSvc) GetArticle(ctx context.Context, id string) *article.Article {
	return &article.Article{}
}

func (svc articleSvc) UpdateArticle(ctx context.Context, dto v1request.UpdateArticleDTO, id string) error {
	return nil
}

func (svc articleSvc) CreateArticle(ctx context.Context, dto v1request.CreateArticleDTO) error {
	return nil
}

func (svc articleSvc) DeleteArticle(ctx context.Context, id string) error {
	return nil
}
