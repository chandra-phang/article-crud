package log

// logging.go provides a simpler way to call LogInfoFields/LogErrorFields,
// while supporting ctx.

import (
	"article-crud/lib"
	"context"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Builder builds up a log entry. Prefer using Infof or Errorf directly.
func Builder(ctx context.Context) *LogBuilder {
	return newBuilder(ctx)
}

func Infof(ctx context.Context, msg string, args ...any) {
	newBuilder(ctx).
		WithSource(lib.WhoCalledMe()).
		Now().Infof(msg, args...)
}

func Errorf(ctx context.Context, err error, msg string, args ...any) {
	newBuilder(ctx).
		WithSource(lib.WhoCalledMe()).
		WithError(err).
		Now().Errorf(msg, args...)
}
