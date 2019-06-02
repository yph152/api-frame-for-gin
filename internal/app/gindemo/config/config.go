package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var (
	global *Config
)

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(fpath string) error {
	c, err := ParseConfig(fpath)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// GetGlobalConfig 获取全局配置
func GetGlobalConfig() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// ParseConfig 解析配置文件
func ParseConfig(fpath string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(fpath, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	RunMode     string      `toml:"run_mode"`
	Swagger     string      `toml:"swagger"`
	Log         Log         `toml:"log"`
	LogGormHook LogGormHook `toml:"log_gorm_hook"`
	HTTP        HTTP        `toml:"http"`
	Monitor     Monitor     `toml:"monitor"`
	CORS        CORS        `toml:"cors"`
	Gorm        Gorm        `toml:"gorm"`
	MySQL       MySQL       `toml:"mysql"`
	Postgres    Postgres    `toml:"postgres"`
}

// Log 日志配置参数
type Log struct {
	Level         int    `toml:"level"`
	Format        string `toml:"format"`
	Output        string `toml:"output"`
	OutputFile    string `toml:"output_file"`
	EnableHook    bool   `toml:"enable_hook"`
	Hook          string `toml:"hook"`
	HookMaxThread int    `toml:"gook_max_thread"`
	HookMaxBuffer int    `toml:"hook_max_buffer"`
}

// LogGormHook 日志gorm钩子配置
type LogGormHook struct {
	DBType       string `toml:"db_type"`
	MaxLifetime  int    `toml:"max_lifetime"`
	maxOpenConn  int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	Table        string `toml:"table"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// Monitor 监控配置参数
type Monitor struct {
	Enable    bool   `toml:"enable"`
	Addr      string `toml:"addr"`
	ConfigDir string `toml:"config_dir"`
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     `toml:"enable"`
	AllowOrigins     []string `toml:"allow_origins"`
	AllowMethods     []string `toml:"allow_methods"`
	AllowHeaders     []string `toml:"allow_headers"`
	AllowCredentials bool     `toml:"allow_credentials"`
	MaxAge           int      `toml:"max_age"`
}

// MySQL mysql 配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)%s?%s", a.User, a.Password, a.Host, a.Parameters)
}

// Postgres postgres配置参数
type Postgres struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
}

// DSN 数据库连接串
func (a Postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", a.Host, a.Port, a.User, a.DBName, a.Password)
}
