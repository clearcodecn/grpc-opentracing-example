package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	logrus.FieldLogger
}

func WithContext(ctx context.Context, module ...string) Logger {
	if len(module) > 0 {
		return logrus.StandardLogger().WithContext(ctx).WithField("module", module[0])
	}
	return logrus.StandardLogger().WithContext(ctx)
}
