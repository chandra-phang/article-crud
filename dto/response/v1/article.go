package v1

import "article-crud/domain/article"

type GetArticleDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type ListArticleDTO struct {
	Articles []GetArticleDTO `json:"articles"`
}

func (dto *GetArticleDTO) ConvertFromArticleEntity(entity *article.Article) *GetArticleDTO {
	return &GetArticleDTO{
		Title:       entity.Title,
		Description: entity.Description,
		Content:     entity.Content,
	}
}

func (dto *ListArticleDTO) ConvertFromArticlesEntity(entities []article.Article) *ListArticleDTO {
	resp := &ListArticleDTO{}
	for _, entity := range entities {
		article := new(GetArticleDTO).ConvertFromArticleEntity(&entity)
		resp.Articles = append(resp.Articles, *article)
	}

	return resp
}
