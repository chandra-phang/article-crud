package v1

type CreateArticleDTO struct {
	Title   string `json:"title" validate:"required"`
	Desc    string `json:"desc"`
	Content string `json:"content" validate:"required"`
}

type UpdateArticleDTO struct {
	Title   string `json:"title" validate:"required"`
	Desc    string `json:"desc"`
	Content string `json:"content" validate:"required"`
}
