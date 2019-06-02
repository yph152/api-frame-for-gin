package apis

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app/options"
	"github.com/yph152/api-frame-for-gin/src/apis/v1alpha1/demo"
)

func InitWeb(ctx context.Context, obj *Object) *gin.Engine {
	gin.SetMode("debug")

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	apiPrefixes := []string{"/apis/"}

	// 跟踪ID
	app.Use(middlerware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	// 访问日志
	app.Use(middlerware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))

	//崩溃恢复
	app.Use(middleware.RecoveryMiddleware())

	// 跨域请求
	if cfg.CORS.Enable {
		app.Use(middleware.CoreMiddleware())
	}

	// 注册/apis 路由
	apis.RegisterRouter(app, obj.Bll)

	// swagger文档
	if dir := cfg.Swagger; dir != "" {
		app.Static("/swagger", dir)
	}

	//静态站点
	if dir := cfg.WWW; dir != "" {
		app.Use(middleware.WWWMiddleware(dir))
	}

	return app
}

func InitHTTPServer(ctx context.Context, handler http.Hanler) func() {
	cfg := config.GetGlobalConfig().HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.StartSpan(ctx, "HTTP服务初始化", "gindemo.InitHTTPServer").Printf("HTTP服务开始启动,地址监听在: [%s]", addr)
		err := server.ListenAndServer()
		if err != nil && err != http.ErrServerClosed {
			logger.StartSpan(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err != server.Shutdown(ctx); err != nil {
			logger.StartSpan(ctx, "关闭HTTP服务", "gindemo.InitHTTPServer").Errorf(err.Error())
		}
	}
}
