package v1

import (
	"article-crud/api/controller"
	"article-crud/api/controller/parse"
	"article-crud/application"
	v1request "article-crud/dto/request/v1"
	v1response "article-crud/dto/response/v1"
	"article-crud/log"
	"encoding/json"
	"errors"
	"io/ioutil"
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

// ListArticles godoc
// @Summary      Get list of articles
// @Description  This API will fetch and return list of articles
// @Tags         Article
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.SuccessResponse{result=v1.ListArticleDTO}
// @Router       /v1/articles [get]
func (c *articleController) ListArticles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articles := c.svc.ListArticles(ctx)

	dto := new(v1response.ListArticleDTO).ConvertFromArticlesEntity(articles)

	controller.WriteSuccess(ctx, w, r, http.StatusOK, dto)
}

// GetArticle godoc
// @Summary      Get article by ID
// @Description  This API will fetch and return article by ID
// @Tags         Article
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.SuccessResponse{result=v1.GetArticleDTO}
// @Failure      400  {object}  response.FailureResponse
// @Failure      404  {object}  response.FailureResponse
// @Router       /v1/articles/id [get]
func (c *articleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// extract path param
	id, err := parse.PathParam(r, "id")

	// If the parameter is not found, a 400 Bad Request error is returned to the client
	if errors.Is(err, parse.ErrParamNotFound) {
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	article, err := c.svc.GetArticle(ctx, id)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][GetArticle] Failed to get article with ID: %s ", id)
		controller.WriteError(ctx, w, r, http.StatusNotFound, err)
		return
	}

	dto := new(v1response.GetArticleDTO).ConvertFromArticleEntity(article)

	controller.WriteSuccess(ctx, w, r, http.StatusOK, dto)
}

// CreateArticle godoc
// @Summary      Create an article
// @Description  This API will create an article
// @Tags         Article
// @Accept       json
// @Produce      json
// @Success      201  {object}  response.SuccessResponse
// @Failure      400  {object}  response.FailureResponse
// @Failure      500  {object}  response.FailureResponse
// @Router       /v1/articles [post]
func (c *articleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody, _ := ioutil.ReadAll(r.Body)
	dto := v1request.CreateArticleDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ArticleController][CreateArticle] Failed to unmarshal request body %v into dto", reqBody)
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	err := dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][CreateArticle] Validation failed for request dto %v ", dto)
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	err = c.svc.CreateArticle(ctx, dto)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][CreateArticle] Failed to create article for request dto %v ", dto)
		controller.WriteError(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	controller.WriteSuccess(ctx, w, r, http.StatusCreated, nil)
}

// UpdateArticle godoc
// @Summary      Update an article
// @Description  This API will update an article
// @Tags         Article
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.SuccessResponse
// @Failure      400  {object}  response.FailureResponse
// @Failure      500  {object}  response.FailureResponse
// @Router       /v1/articles/id [put]
func (c *articleController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// extract path param
	id, err := parse.PathParam(r, "id")

	// If the parameter is not found, a 400 Bad Request error is returned to the client
	if errors.Is(err, parse.ErrParamNotFound) {
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	dto := v1request.UpdateArticleDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ArticleController][UpdateArticle] Failed to unmarshal request body %v into dto", reqBody)
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	err = dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][UpdateArticle] Validation failed for request dto %v ", dto)
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	err = c.svc.UpdateArticle(ctx, dto, id)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][UpdateArticle] Failed to update article for request dto %v ", dto)
		controller.WriteError(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	controller.WriteSuccess(ctx, w, r, http.StatusOK, nil)
}

// DeleteArticle godoc
// @Summary      Delete an article
// @Description  This API will delete an article
// @Tags         Article
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.SuccessResponse
// @Failure      400  {object}  response.FailureResponse
// @Failure      500  {object}  response.FailureResponse
// @Router       /v1/articles/id [delete]
func (c *articleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// extract path param
	id, err := parse.PathParam(r, "id")

	// If the parameter is not found, a 400 Bad Request error is returned to the client
	if errors.Is(err, parse.ErrParamNotFound) {
		controller.WriteError(ctx, w, r, http.StatusBadRequest, err)
		return
	}

	err = c.svc.DeleteArticle(ctx, id)
	if err != nil {
		log.Errorf(ctx, err, "[ArticleController][DeleteArticle] Failed to delete article for id %s ", id)
		controller.WriteError(ctx, w, r, http.StatusInternalServerError, err)
		return
	}

	controller.WriteSuccess(ctx, w, r, http.StatusOK, nil)
}
