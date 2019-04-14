package app

import (
	"github.com/golang/glog"

	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app/options"
	"github.com/yph152/api-frame-for-gin/src/apis"
)

func Run(opt *options.ServerRunOptions, stopCh <-chan int) error {
	err := apis.Server(opt)
	if err != nil {
		glog.Errorf("Server Run failed: %v", err.Error())
		return err
	}
	return nil
}
