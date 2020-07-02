# 如何利用 Go kit 编写微服务

## 基础的实现
参考代码 stringsvc1

依赖需求关系：Transport-->Endpoint-->Service

### 如何编写一个 Go kit 应用
1. 确定业务的 service interface
2. 实现 interface
3. 将 service 集成到 endpoint 中
4. 创建 transport

### 如何创建 endpoint
在 Go kit 的设计哲学中，Endpoint 就像是一个 adapter。对外，Endpoint 接收 request，对内 Endpoint 调用 Service 的相应 method 完成业务逻辑，将 method 的 response 返回。
在 Go kit 内部，endpoint 的定义如下：

```
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
```

所以，endpoint 实际上就是一个函数，它接收一个 context 和 request，返回 response 和 error。那么，如果我们要实现一个自己的 endpoint，只需要实现一个同样的函数，在函数的内部，通过调用 service 来完成业务逻辑即可。

```
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}
```
