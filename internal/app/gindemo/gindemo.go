package gindemo

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app/options"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/bll"
	//"github.com/yph152/api-frame-for-gin/internal/app/gindemo/config"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/model"
	"github.com/yph152/api-frame-for-gin/pkg/logger"
	"github.com/yph152/api-frame-for-gin/pkg/util"
)

// Object 对象集合
type Object struct {
	Model *model.Common
	Bll   *bll.Common
}

// Init 应用初始化
func Init(ctx context.Context) func() {
	loggerCall, err := InitLogger()
	if err != nil {
		panic(err)
	}

	obj, objCall, err := InitObject(ctx)
	if err != nil {
		panic(err)
	}

	app := InitWeb(ctx, obj)
	httpCall := InitHTTPServer(ctx, app)

	return func() {
		if httpCall != nil {
			httpCall()
		}

		if objCall != nil {
			objCall()
		}

		if loggerCall != nil {
			loggerCall()
		}
	}
}

// InitObject 初始化对象数据
func InitObject(ctx context.Context) (*Object, func(), error) {
	obj := &Object{}

	return obj, func() {
	}, nil
}

func Server(opt *options.ServerRunOptions) error {
	//cfg := config.LoadGlobalConfig(opt.Config)
	//swaggerDir := opt.SwaggerDir

	ctx := logger.NewTraceIDContext(context.Background(), util.MustUUID())
	span := logger.StartSpanWithCall(ctx, "主函数", "main")
	span().Printf("服务启动, 运行模式: %s, 版本号: %s, 进程号: %d", "debug", "1.0", os.Getpid())

	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	call := Init(ctx)
	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		span().Printf("获取到退出信号[%s]", sig.String())
	}

	if call != nil {
		call()
	}
	span().Printf("服务退出")

	return nil
}
