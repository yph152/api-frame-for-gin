package options

import (
	"github.com/spf13/pflag"
)

// 命令行参数列表
type ServerRunOptions struct {
	Port       int
	Config     string
	SwaggerDir string
}

func NewServerRunOptions() *ServerRunOptions {
	opt := &ServerRunOptions{}
	return opt
}

// 添加命令行参数
func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.Port, "port", 8080, "api port")
	fs.StringVar(&s.Config, "config", "/tmp/config.toml", "server config.")
	fs.StringVar(&s.SwaggerDir, "swagger_dir", "/tmp/swagger", " swagger direction.")
}
