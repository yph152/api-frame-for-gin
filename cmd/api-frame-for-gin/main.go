package main

import (
	"github.com/golang/glog"
	"github.com/spf13/pflag"

	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app"
	"github.com/yph152/api-frame-for-gin/cmd/api-frame-for-gin/app/options"
)

func main() {
	var stopCh <-chan int

	s := options.NewServerRunOptions()
	s.AddFlags(pflag.CommandLine)

	pflag.Parse()
	defer glog.Flush()
	if err := app.Run(s, stopCh); err != nil {
		glog.Error("app exit...")
	}
}
