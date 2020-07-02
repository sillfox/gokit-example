## 模块拆分
在基于 stringsvc1 的基础上，将代码按照模块拆分出去。
将以下部分，拆分到 service.go
* type StringService
* type stringService
    * func Uppercase
    * func Count
* var ErrEmpty

将以下部分，拆分到 transport.go
* func makeUppercaseEndpoint
* func makeCountEndpoint
* func decodeUppercaseRequest
* func decodeCountRequest
* func encodeResponse
* type uppercaseRequest
* type uppercaseResponse
* type countRequest
* type countResponse

>  两个makeXXXEndpoint 也是属于 transport 的范畴。

## 添加 Log Middleware

在 Go kit 内部，Middleware 被定义为：

```
type Middleware func(Endpoint) Endpoint
```

它是一个函数，接收 Endpoint，返回 Endpoint，就是一个 Endpoint 的 decorator。
同样地，我们也可以使用类似于 Endpoint 封装 Service 的这个方式，用 Middleware 将 Endpoint 封装起来。
这样，我们就创建了一个基于 Endpoint Middleware。它可以在 Endpoint 之外实现一些与 Endpoint 层相关的其他逻辑。但是，有时候我们会需要一个与业务数据相关的 Middleware，这就需要一个基于 Service 的 Middleware 了。接下来我们会看到如果创建一个基于 Service 的 Middleware。