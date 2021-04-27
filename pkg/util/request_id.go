package util

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	RequestIDKey = "request_id"
)

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			req := ectx.Request()
			res := ectx.Response()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = uuid.New().String()
			}
			// echoのヘッダーとContextにRequestIDを設定
			res.Header().Set(echo.HeaderXRequestID, rid)
			ectx.Set(RequestIDKey, rid)
			return next(ectx)
		}
	}
}

func WithRequestID(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestId)
}

func GetRequestID(ctx context.Context) string {
	if requestID, exists := ctx.Value(RequestIDKey).(string); exists {
		return requestID
	}
	return "request id not found"
}
