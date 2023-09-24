// Package parse provides parsing of URL and Query params
package parse

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var ErrParamNotFound = errors.New("param not found")
var ErrParamInvalidType = errors.New("invalid type for param")

type TypeConv[T any] func(s string) (T, error)

var (
	AsInt   TypeConv[int]   = strconv.Atoi
	AsInt64 TypeConv[int64] = func(s string) (int64, error) { return strconv.ParseInt(s, 10, 64) }
)

func newParamNotFound(key string) error {
	return fmt.Errorf("%w: %s", ErrParamNotFound, key)
}

func newParamInvalidType(key string, reason error) error {
	return fmt.Errorf("%w: %s (%v)", ErrParamInvalidType, key, reason)
}

// PathParam returns a path parameter as a string type.
// Returns ErrParamNotFound if
func PathParam(r *http.Request, key string) (string, error) {
	v := chi.URLParam(r, key)
	if len(v) == 0 {
		return "", newParamNotFound(key)
	}
	return v, nil
}

func QueryParams(r *http.Request, key string) Params[string] {
	values := r.URL.Query()[key]
	asString := func(s string) (string, error) { return s, nil }
	return NewParams(key, values, asString)
}

func QueryParamsType[T any](r *http.Request, key string, typeconv TypeConv[T]) Params[T] {
	values := r.URL.Query()[key]
	return NewParams(key, values, typeconv)
}
