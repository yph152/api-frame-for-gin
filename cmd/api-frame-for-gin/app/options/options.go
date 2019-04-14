package options

import (
	"github.com/spf13/pflag"
)

type ServerRunOptions struct {
	Port int
}

func NewServerRunOptions() *ServerRunOptions {
	opt := &ServerRunOptions{}
	return opt
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.IntVar(&s.Port, "port", 8080, "api port")
}
