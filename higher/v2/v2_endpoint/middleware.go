package v2_endpoint

import (
	"apple/higher/v2/v2_service"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(v2_service.ContextReqUUid)), zap.Any("调用 v2_endpoint LoggingMiddleware", "处理完请求"), zap.Any("耗时毫秒", time.Since(begin).Milliseconds()))
			}(time.Now())
			return next(ctx, request)
		}
	}
}
