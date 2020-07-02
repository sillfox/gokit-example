## 基于 Service 的 Middleware

准确来说，基于 Service 的 Middleware，其实并不是 Middleware。因为它不属于 Middleware 类型，我们只是借用了与 Middleware 类似的 decorator 模式，将 service 包裹在一个 Middleware 中。

回到我们的 StringService interface。试想，如果我实现一个 StringService 的 struct，其内部包裹着 stringService，所有 interface定义的 method 都通过调用 stringService 来完成，而这个代理 struct 只需要在调用 stringService method 的前后完成日志的打印。这样是不是也实现了同样的需求呢？

代码如下：

```
type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.Count(s)
	return
}
```

> 因为这个 Middleware 关心的是与 Service 相关的东西，所以，如果再使用之前版本的基于 Endpoint 的 Middleware 的话，就会导致依赖关系混乱。在 Endpoint 层处理 Service 相关的逻辑。所以，如果需要处理 Service 相关的逻辑的时候，就应该使用基于 Service 的 decorator

## main func
再回到 main 函数中，因为我们实现的 Middleware 实际上是一个 service，所以，我们还是需要 makeXXXEndpoint 函数，将我们的 Middleware service 传进去生成一个 endpoint 实例。