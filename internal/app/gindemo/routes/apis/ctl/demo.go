package demo

import (
	"github.com/gin-gonic/gin"
)

type Demo struct {
}

func NewDemo() Demo {
	return Demo{}
}

func (d Demo) GetDemo(c *gin.Context) {
	c.String(200, "", "demo")
}
