package main

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// proxymw 实现了 StringService 接口。将 uppercase 请求转发给了 endpoint。
// 同时将其他的请求都交给了 next 来处理
type proxymw struct {
	next      StringService     // 直接调用的 service
	uppercase endpoint.Endpoint // 通过 endpoint 调用外部的 service
}

func (mw proxymw) Uppercase(s string) (string, error) {
	response, err := mw.uppercase(uppercaseRequest{S: s})
	if err != nil {
		return "", err
	}
	resp := response.(uppercaseResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func (mw proxymw) Count(s string) int {
	return mw.next.Count(s)
}

func makeUppercaseProxy(proxyURL string) endpoint.Endpoint {
	return httptransport.NewClient(
		"GET",
		mustParsseURL(proxyURL),
		encodeUppercaseRequest,
		decodeUppercaseResponse,
	).Endpoint()
}
