package conf

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
)

type RedisConf struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}
type Config struct {
	Debug bool      `toml:"debug"`
	Port  int64     `toml:"port"`
	DBDsn string    `toml:"db_dsn"`
	Redis RedisConf `toml:"redis"`
}

func NewConfig() *Config {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	work, err := os.Open(wd + "/config.toml")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(work)
	if err != nil {
		panic(err)
	}
	var conf Config
	err = toml.Unmarshal(body, &conf)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	return &conf
}
