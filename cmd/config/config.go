package config

import (
	"flag"
	"github.com/caarlos0/env"
)

type Config struct {
	Host string `env:"RUN_ADDR"`
}

func NewConfig() *Config {

	conf := Config{}

	flag.StringVar(&conf.Host, "a", ":8080", "address and port to run server")
	flag.Parse()

	err := env.Parse(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
