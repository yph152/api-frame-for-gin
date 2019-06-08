package ctl

import (
	"github.com/gin-gonic/gin"

	"github.com/yph152/api-frame-for-gin/internal/app/gindemo/bll"
)

type Demo struct {
	DemoBll *bll.Demo
}

func NewDemo(b *bll.Common) *Demo {
	return &Demo{
		DemoBll: b.Demo,
	}
}

func (a *Demo) Query(c *gin.Context) {
	c.JSON(200, nil)
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
