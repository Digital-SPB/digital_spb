package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  `yaml:"app"`
	HTTP `yaml:"http"`
	Log  `yaml:"log"`
	PG   `yaml:"postgres"`
}

type App struct {
	Name    string `yaml:"name" env-required:"true" env:"APP_NAME"`
	Version string `yaml:"version" env-required:"true" env:"APP_VERSION"`
}

type HTTP struct {
	Port string `yaml:"port" env-required:"true" env:"HTTP_PORT"`
}

type Log struct {
	Level string `yaml:"level" env-required:"true" env:"LOG_LEVEL"`
}
type PG struct {
	MaxPoolSize int    `yaml:"max_pool_size" env-required:"true" env:"PG_MAX_POOL_SIZE"`
	URL         string `yaml:"url" env-required:"true" env:"PG_URL"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env vars: %w", err)
	}

	return cfg, nil
}
