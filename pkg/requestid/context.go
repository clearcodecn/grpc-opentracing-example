package requestid

import (
	"context"
	"github.com/google/uuid"
)

const (
	key = "x-request-id"
)

func NewWithContext(ctx context.Context, id string) context.Context {
	if id == "" {
		id = uuid.New().String()
	}
	return context.WithValue(ctx, key, id)
}

func RequestID(ctx context.Context) string {
	reqid, _ := ctx.Value(key).(string)
	return reqid
}
