package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	// loggingMiddleware struct version

	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	uppcaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppcaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}
