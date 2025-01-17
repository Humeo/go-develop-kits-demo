// main.go
package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tracer trace.Tracer
	logger *zap.Logger
)

// 初始化 tracer
func initTracer() (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	// 创建 OTLP exporter
	conn, err := grpc.Dial(
		"localhost:4317",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create exporter: %w", err)
	}

	// 创建资源属性
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("demo-service2"),
			semconv.ServiceVersionKey.String("1.0.0"),
			attribute.String("environment", "development"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 创建 TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)

	// 设置全局 TracerProvider
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer("demo-tracer")

	return tp, nil
}

// 初始化 logger
func initLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	return config.Build()
}

// HTTP 中间件
func tracingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 创建 span
		ctx, span := tracer.Start(ctx, "http_request",
			trace.WithAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.url", r.URL.String()),
				attribute.String("http.user_agent", r.UserAgent()),
			),
		)
		defer span.End()

		// 使用带有追踪信息的 logger
		requestLogger := logger.With(
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("span_id", span.SpanContext().SpanID().String()),
		)

		requestLogger.Info("Received HTTP request")

		// 调用下一个处理器
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// 模拟数据库操作
func simulateDBOperation(ctx context.Context) error {
	_, span := tracer.Start(ctx, "database_query",
		trace.WithAttributes(
			attribute.String("db.system", "postgresql"),
			attribute.String("db.operation", "select"),
		),
	)
	defer span.End()

	// 模拟数据库延迟
	time.Sleep(100 * time.Millisecond)

	return nil
}

// 处理首页请求
func handleHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)

	requestLogger := logger.With(
		zap.String("trace_id", span.SpanContext().TraceID().String()),
		zap.String("span_id", span.SpanContext().SpanID().String()),
	)

	// 模拟数据库操作
	if err := simulateDBOperation(ctx); err != nil {
		requestLogger.Error("Database operation failed", zap.Error(err))
		span.RecordError(err)
		span.SetStatus(codes.Error, "Database operation failed")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	requestLogger.Info("Request processed successfully")
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// 初始化 logger
	var err error
	logger, err = initLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// 初始化 tracer
	tp, err := initTracer()
	if err != nil {
		logger.Fatal("Failed to initialize tracer", zap.Error(err))
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error("Error shutting down tracer provider", zap.Error(err))
		}
	}()

	// 设置路由
	http.HandleFunc("/", tracingMiddleware(handleHome))

	// 启动服务器
	logger.Info("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
