package main

import (
	"fmt"
	"net/http"
	"time"
)

// 中间件类型定义
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 增强的中间件链接器
func chainMiddleware(middlewares ...Middleware) Middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 从最后一个中间件开始，逐个包装
			last := final
			for i := len(middlewares) - 1; i >= 0; i-- {
				last = middlewares[i](last)
			}
			// 执行完整的中间件链
			last(w, r)
		}
	}
}

// 跟踪中间件
func tracingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceID := generateTraceID()
		ctx := context.WithValue(r.Context(), "trace_id", traceID)

		fmt.Printf("[Trace:%s] Request started\n", traceID)
		start := time.Now()

		next.ServeHTTP(w, r.WithContext(ctx))

		fmt.Printf("[Trace:%s] Request completed in %v\n",
			traceID, time.Since(start))
	}
}

// 认证中间件
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID := validateToken(token)
		ctx := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// 日志中间件
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取之前中间件设置的值
		traceID := r.Context().Value("trace_id").(string)
		userID := r.Context().Value("user_id").(string)

		fmt.Printf("[%s] User:%s accessing %s\n",
			traceID, userID, r.URL.Path)

		next.ServeHTTP(w, r)
	}
}

// 实际使用示例
func main() {
	// 1. 基本用法
	handler := chainMiddleware(
		tracingMiddleware,
		authMiddleware,
		loggingMiddleware,
	)(finalHandler)

	// 2. 条件中间件
	handler = chainMiddleware(
		tracingMiddleware,
		// 根据条件选择是否包含某个中间件
		func(next http.HandlerFunc) http.HandlerFunc {
			if enableAuth {
				return authMiddleware(next)
			}
			return next
		},
		loggingMiddleware,
	)(finalHandler)

	// 3. 动态中间件列表
	middlewares := []Middleware{tracingMiddleware}
	if enableAuth {
		middlewares = append(middlewares, authMiddleware)
	}
	if enableLogging {
		middlewares = append(middlewares, loggingMiddleware)
	}

	handler = chainMiddleware(middlewares...)(finalHandler)

	http.HandleFunc("/", handler)
}
