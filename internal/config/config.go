package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env          string        `envconfig:"ENV" default:"development" required:"true" desc:"development, test or production"`
	Port         uint          `envconfig:"PORT" default:"5000" required:"true" desc:"HTTP port to listen"`
	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s" desc:"tcp connection read timeout"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s" desc:"tcp connection write timeout"`
	IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s" desc:"tcp connection idle timeout"`
}

// FromEnv load configuration from env vars
func FromEnv() (Config, error) {
	c := Config{}
	return c, envconfig.Process("", &c)
}
