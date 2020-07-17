package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

// middleware
// go-kit内部，Middleware 是一个函数类型：type Middleware func(Endpoint) Endpoint
// 想要使用自己的 middleware，需要像下面这样定义：

func loggingMiddleware(logger log.Logger) endpoint.Middleware { // 创建一个middleware
	return func(next endpoint.Endpoint) endpoint.Endpoint { // middleware 函数的形参和返回值都是 endpoint
		return func(ctx context.Context, request interface{}) (interface{}, error) { // endpoint
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request) // 调用内层的 endpoint
		}
	}
}

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	// logginMiddleware func version
	// go-kit 并没有提供 IoC 来实现依赖注入，而是倡导自己管理依赖
	svc := stringService{}
	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	var count endpoint.Endpoint
	count = makeCountEndpoint(svc)
	count = loggingMiddleware(log.With(logger, "method", "count"))(count)
	uppercaseHandler := httptransport.NewServer(
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)
	countHandler := httptransport.NewServer(
		count,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
