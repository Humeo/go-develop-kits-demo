package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 请求前记录时间
		startTime := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// 调用下一个处理函数
		next(w, r)

		// 请求后记录耗时
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(startTime))
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从header中获取token
		token := r.Header.Get("Authorization")

		// 这里简单判断token是否为"valid-token"
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 认证通过，调用下一个处理函数
		next(w, r)
	}
}

// 实际的业务处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// 组合多个中间件
func Chain(f http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func main() {
	// 组合中间件和处理函数
	handler := Chain(helloHandler, LoggerMiddleware, AuthMiddleware)

	// 注册路由
	http.HandleFunc("/hello", handler)

	// 启动服务器
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
