package ctl

import (
	"github.com/gin-gonic/gin"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/bll"
	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/schema"
)

// @Name Demo
// @Description demo
type Demo struct {
	DemoBll *bll.Demo
}

// NewDemo 创建demo控制器
func NewDemo(b *bll.Common) *Demo {
	return &Demo{
		DemoBll: b.Demo,
	}
}

// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param code query string false "编号"
// @Param name query string false "名称"
// @Success 200 []schema.Demo "查询结果: {list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
func (a *Demo) Query(c *gin.Context) {
	c.JSON(200, []schema.Demo{})
}

func (a *Demo) QueryPage(c *gin.Context) {
	c.JSON(200, nil)
}

func (a *Demo) Get(c *gin.Context) {
	c.JSON(200, nil)
}

func (a *Demo) Create(c *gin.Context) {
	c.JSON(200, nil)
}

func (a *Demo) Update(c *gin.Context) {
	c.JSON(200, nil)
}

func (a *Demo) Delete(c *gin.Context) {
	c.JSON(200, nil)
}
