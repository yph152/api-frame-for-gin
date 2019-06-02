package apis

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app/options"
	"github.com/yph152/api-frame-for-gin/src/apis/v1alpha1/demo"
)

func Server(opt *options.ServerRunOptions) error {

	d := demo.NewDemo()

	r := gin.Default()
	r.GET("/apis/v1alpha1/demo", d.GetDemo)

	r.Run(fmt.Sprintf(":%v", strconv.Itoa(opt.Port)))
	return nil
}
