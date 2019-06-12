package gindemo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/config"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/middleware"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/routers/apis"
	"github.com/yph152/api-frame-for-gin/pkg/logger"
)

func InitWeb(ctx context.Context, obj *Object) *gin.Engine {
	gin.SetMode("debug")

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	apiPrefixes := []string{"/apis/"}

	// 跟踪ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 访问日志
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	//崩溃恢复
	app.Use(middleware.RecoveryMiddleware())

	app.Static("/swagger", "../../internal/app/gindemo/swagger")
	cfg := config.GetGlobalConfig()
	// 跨域请求
	if cfg.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// 注册/apis 路由
	apis.RegisterRoute(app, obj.Bll)

	// swagger文档
	if dir := cfg.Swagger; dir != "" {
		app.Static("/swagger", dir)
	}

	return app
}

// http 服务初始化
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	//var err error
	cfg := config.GetGlobalConfig().HTTP
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", 8088)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.StartSpan(ctx, "HTTP服务初始化", "gindemo.InitHTTPServer").Printf("HTTP服务开始启动,地址监听在: [%s]", addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			//logger.StartSpan(ctx, "服务初始化失败", time.Second*time.Duration(cfg.ShutdownTimeout))
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.StartSpan(ctx, "关闭HTTP服务", "gindemo.InitHTTPServer").Errorf(err.Error())
		}
	}
}
