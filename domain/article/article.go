package article

import "context"

type Article struct {
	ID          string
	Title       string
	Description string
	Content     string
}

type IArticleRepository interface {
	FindArticles(ctx context.Context, offset int64, limit int64) ([]*Article, error)
	FindArticleByURI(ctx context.Context, uri string) (*Article, error)
	Save(ctx context.Context, segmentEntity *Article) error
}
