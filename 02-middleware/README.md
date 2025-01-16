## 中间件是什么？
中间件是一种软件设计模式。主要用于在应用程序的请求-响应周期中添加额外的处理层。
它能帮助我们实现关注点分离，使代码更加模块化和可维护。

- 关注点分离(Separation of Concerns, SoC)是一个重要的软件设计原则，它强调将程序分解为不同的部分，每个部分负责处理特定的功能或问题。 
  - 基本概念
  1. 将不同的功能模块相互分离
  2. 每个模块只关注自己需要处理的特定问题
  3. 减少模块之间的耦合度
  4. 提高代码的可维护性和复用性

1. 基本概念
中间件是位于应用程序核心逻辑之外的一层处理程序
可以在请求到达实际处理程序之前或之后执行
通常采用链式处理方式（中间件链）
2. 常见用途
* 日志记录
* 身份认证
* 权限验证
* 请求参数验证
* 跨域(CORS)处理
* 性能监控
* 错误处理

3. 中间件的特点
* 可重用性：同一个中间件可以用于多个路由或处理器
* 可组合性：多个中间件可以组合使用
* 顺序性：中间件的执行顺序很重要
4. 常见框架中的中间件
* Gin框架：使用 gin.Use(middleware)
* Echo框架：使用 e.Use(middleware)
* 标准库：使用 http.HandleFunc 包装

5. 最佳实践
* 中间件应该只负责单一职责
* 注意中间件的执行顺序
* 避免在中间件中执行太重的操作
* 正确处理错误和异常情况

## How to use demo?
### net_http
使用`net/http`实现的中间件处理流程
* S1: 运行 net/http/main.go
* S2: 请求

windows:
```ps
curl.exe -H "Authorization: valid-token" http://localhost:8080/hello
```
linux:
```shell
curl -H "Authorization: valid-token" http://localhost:8080/hello
```
执行结果
服务端：
```
Server starting on port 8080...
2025/01/14 20:23:25 Started GET /hello
2025/01/14 20:23:25 Completed GET /hello in 25.0866ms

```
请求端:
```
Hello, World!
```

中间件代码解读：
```go
func Chain(
    f http.HandlerFunc,  // 最终的处理函数
    middlewares ...func(http.HandlerFunc) http.HandlerFunc,  // 可变数量的中间件函数
) http.HandlerFunc {
    // 从左到右遍历所有中间件
    for _, m := range middlewares {
        f = m(f)  // 将当前处理函数包装到中间件中
    }
    return f  // 返回包装后的处理函数
}
```
* 第一个参数是最终的处理函数
f http.HandlerFunc

* 第二个参数是可变数量的中间件函数 
每个中间件函数接收一个 http.HandlerFunc 并返回一个新的 http.HandlerFunc
middlewares ...func(http.HandlerFunc) http.HandlerFunc

```
// 假设我们这样调用：
handler := Chain(helloHandler, LoggerMiddleware, AuthMiddleware)

// 执行顺序是：
// 1. f = LoggerMiddleware(helloHandler)
// 2. f = AuthMiddleware(LoggerMiddleware(helloHandler))
```

请求处理流程
```
// 当请求进来时，执行顺序是：
AuthMiddleware
    -> LoggerMiddleware
        -> helloHandler
        <- LoggerMiddleware
    <- AuthMiddleware

```

让我们用一个完整的示例来演示：

```go
package main

import (
    "fmt"
    "net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("1. Logger 开始")
        next(w, r)
        fmt.Println("4. Logger 结束")
    }
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("0. Auth 开始")
        next(w, r)
        fmt.Println("5. Auth 结束")
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("2. 进入 Hello 处理函数")
    fmt.Fprintf(w, "Hello, World!")
    fmt.Println("3. 离开 Hello 处理函数")
}

func Chain(f http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}

func main() {
    handler := Chain(helloHandler, LoggerMiddleware, AuthMiddleware)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```

请求到达时
```
0. Auth 开始
1. Logger 开始
2. 进入 Hello 处理函数
3. 离开 Hello 处理函数
4. Logger 结束
5. Auth 结束
```

这种执行顺序的原因是：

1. 中间件像洋葱层一样层层包裹
2. 每个中间件都有"前置"和"后置"操作
3. 执行顺序是从外到内，然后再从内到外
4. 这种结构允许中间件在请求处理前后都能执行代码


这种模式的好处是：
1. 可以在请求处理前进行预处理（如认证、日志记录开始）
2. 可以在请求处理后进行后处理（如记录响应时间、清理资源）
3. 中间件之间相互独立，易于维护和组合
4. 这就像一个洋葱模型：请求必须穿过所有的层才能到达核心（处理函数），然后响应又要穿过所有的层才能返回给客户端。
