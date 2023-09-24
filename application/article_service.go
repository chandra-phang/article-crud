package application

import (
	"article-crud/domain/article"
	v1request "article-crud/dto/request/v1"
	"article-crud/handlers"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type IArticleService interface {
	// svc CRUD methods for domain objects
	ListArticles(ctx context.Context) []article.Article
	GetArticle(ctx context.Context, id string) (*article.Article, error)
	CreateArticle(ctx context.Context, dto v1request.CreateArticleDTO) error
	UpdateArticle(ctx context.Context, dto v1request.UpdateArticleDTO, id string) error
	DeleteArticle(ctx context.Context, id string) error
}

type articleSvc struct {
	Handler handlers.Handler
}

var articleSvcSingleton IArticleService

func InitArticleService(h handlers.Handler) {
	articleSvcSingleton = articleSvc{
		Handler: h,
	}
}

func GetArticleService() IArticleService {
	return articleSvcSingleton
}

func (svc articleSvc) ListArticles(ctx context.Context) []article.Article {
	results, err := svc.Handler.DB.Query("SELECT id,title,description,content,created_at,updated_at FROM articles;")
	if err != nil {
		log.Println("failed to execute query", err)
		return nil
	}

	var articles = make([]article.Article, 0)
	for results.Next() {
		var article article.Article
		err = results.Scan(&article.ID, &article.Title, &article.Description, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil
		}

		articles = append(articles, article)
	}

	return articles
}

func (svc articleSvc) GetArticle(ctx context.Context, id string) (*article.Article, error) {
	queryStmt := `SELECT id,title,description,content,created_at,updated_at FROM articles WHERE id = $1 ;`
	results, err := svc.Handler.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		return nil, err
	}

	var article article.Article
	for results.Next() {
		err = results.Scan(&article.ID, &article.Title, &article.Description, &article.Content, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			log.Println("failed to scan", err)
			return nil, err
		}
	}

	if article.ID == "" {
		return nil, errors.New("article not found")
	}

	return &article, nil
}

func (svc articleSvc) CreateArticle(ctx context.Context, dto v1request.CreateArticleDTO) error {
	article := article.Article{}
	article.ID = (uuid.New()).String()

	query := `INSERT INTO articles (id,title,description,content,created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	err := svc.Handler.DB.QueryRow(query, &article.ID, &dto.Title, &dto.Description, &dto.Content, time.Now(), time.Now()).Scan(&article.ID)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}

func (svc articleSvc) UpdateArticle(ctx context.Context, dto v1request.UpdateArticleDTO, id string) error {
	queryStmt := `UPDATE articles SET title = $2, description = $3, content = $4 WHERE id = $1 RETURNING id;`
	err := svc.Handler.DB.QueryRow(queryStmt, &id, &dto.Title, &dto.Description, &dto.Content).Scan(&id)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}

func (svc articleSvc) DeleteArticle(ctx context.Context, id string) error {
	queryStmt := `DELETE FROM articles WHERE id = $1;`
	_, err := svc.Handler.DB.Query(queryStmt, &id)
	if err != nil {
		log.Println("failed to execute query", err)
		return err
	}

	return nil
}
