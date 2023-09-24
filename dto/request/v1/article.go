package v1

import (
	"article-crud/apperrors"
	"article-crud/log"
	"context"
	"errors"

	"github.com/go-playground/validator"
)

type CreateArticleDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Content     string `json:"content" validate:"required"`
}

type UpdateArticleDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Content     string `json:"content" validate:"required"`
}

func (dto CreateArticleDTO) Validate(ctx context.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		log.Infof(ctx, "[CreateArticleDTO] [Validate] Request DTO validation failed %v",
			map[string]interface{}{
				"error":   vErr,
				"request": dto,
			})
		return errors.New(vErr)
	}

	return nil
}

func (dto UpdateArticleDTO) Validate(ctx context.Context) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		vErr := apperrors.TryTranslateValidationErrors(err)
		log.Infof(ctx, "[UpdateArticleDTO] [Validate] Request DTO validation failed %v",
			map[string]interface{}{
				"error":   vErr,
				"request": dto,
			})
		return errors.New(vErr)
	}

	return nil
}
