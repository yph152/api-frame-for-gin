package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/bll"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/routers/apis/ctl"
)

// RegisterRoute 注册/apis路由
func RegisterRoute(app *gin.Engine, b *bll.Common) {
	g := app.Group("/apis")

	demoCtl := ctl.NewDemo(b)

	v1 := g.Group("/v1")
	{
		v1.GET("/demos", demoCtl.Query)
		v1.GET("/demos/:id", demoCtl.Get)
		v1.POST("/demos", demoCtl.Create)
		v1.PUT("/demos/:id", demoCtl.Update)
		v1.DELETE("/demos/:id", demoCtl.Delete)
	}
}
