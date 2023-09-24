package controller

import (
	"article-crud/apiconst"
	"article-crud/dto/response"
	"context"
	"encoding/json"
	"net/http"

	"log"
)

func WriteSuccess(ctx context.Context, w http.ResponseWriter, r *http.Request, statusCode int, result interface{}) {
	// TODO: implements checks for client connectios
	err := writeJSON(w, r, statusCode, response.SuccessResponse{
		Success: true,
		Result:  result,
	})
	if err != nil {
		log.Print(ctx, err, "error while writing JSON response")
	}
}

func WriteError(ctx context.Context, w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	err = writeJSON(w, r, statusCode, response.FailureResponse{
		Success: false,
		Failure: err,
	})
	if err != nil {
		log.Print(ctx, err, "error while writing JSON response")
	}
}

func writeJSON(w http.ResponseWriter, r *http.Request, statusCode int, response interface{}) error {
	w.Header().Set(apiconst.ContentTypeHeader, apiconst.ContentTypeJSON)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
