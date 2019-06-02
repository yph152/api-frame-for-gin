package gindemo

import "fmt"

// Object 对象集合
type Object struct {
	Model *model.Common
	Bll   *model.Common
}

// Init 应用初始化
func Init(ctx context.Context) func() {
	loggerCall, err := InitLogger()
	if err != nil {
		panic(err)
	}

	if c := config.GetGlobalConfig().Monitor; c.Enable {
		err = agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: true})
		if err != nil {
			logger.StartSpan(ctx, "开启[agent]服务监听", "gindemo.Init").Error(err.Error())
		}
	}

	obj, objCall, err := InitObject()
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
