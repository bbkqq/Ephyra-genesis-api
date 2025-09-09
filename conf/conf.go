package conf

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

//go:embed dev/conf.yaml
//go:embed prod/conf.yaml
var f embed.FS

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz    Hertz       `yaml:"hertz"`
	MySQL    MySQL       `yaml:"mysql"`
	Redis    Redis       `yaml:"redis"`
	Contract []*Contract `yaml:"contract"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	DB       int    `yaml:"db"`
}

type Hertz struct {
	Service         string `yaml:"service"`
	Address         string `yaml:"address"`
	EnablePprof     bool   `yaml:"enable_pprof"`
	EnableGzip      bool   `yaml:"enable_gzip"`
	EnableAccessLog bool   `yaml:"enable_access_log"`
	LogLevel        string `yaml:"log_level"`
	LogFileName     string `yaml:"log_file_name"`
	LogMaxSize      int    `yaml:"log_max_size"`
	LogMaxBackups   int    `yaml:"log_max_backups"`
	LogMaxAge       int    `yaml:"log_max_age"`
}
type Contract struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	RpcUrl  string `yaml:"rpc_url"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	//prefix := "conf"
	confFileRelPath := filepath.Join(filepath.Join(GetEnv(), "conf.yaml"))
	content, err := f.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}

	conf.Env = GetEnv()

	//pretty.Printf("%+v\n", conf)
}

func GetEnv() string {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		return "dev"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
