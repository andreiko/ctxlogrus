package ctxlog

import (
	"context"
	"github.com/Sirupsen/logrus"
)

const logFieldsContextKey = "log_fields"

func GetContextualLogger(context context.Context, logger logrus.FieldLogger) logrus.FieldLogger {
	ctxValue := context.Value(logFieldsContextKey)
	if ctxValue == nil {
		return logger
	}

	ctxFields, ok := ctxValue.(logrus.Fields)
	if !ok {
		return logger
	}

	return logger.WithFields(ctxFields)
}

func GetUpdatedLoggingContext(ctx context.Context, logger logrus.FieldLogger, fields logrus.Fields) (context.Context, logrus.FieldLogger) {
	newFields := make(logrus.Fields)

	ctxValue := ctx.Value(logFieldsContextKey)
	if ctxValue != nil {
		if currentFields, ok := ctxValue.(logrus.Fields); ok {
			for name, value := range currentFields {
				newFields[name] = value
			}
		}
	}

	for name, value := range fields {
		newFields[name] = value
	}

	return context.WithValue(ctx, logFieldsContextKey, newFields), logger.WithFields(newFields)
}
