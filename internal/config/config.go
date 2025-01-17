package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service Service `yaml:"service"`
	Tts     Tts     `yaml:"tts"`
	Minio   Minio   `yaml:"minio"`
}

var instance *Config

func Init(pathConfig string) *Config {
	slog.Info("read application config")
	instance = &Config{}
	if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
		slog.Error(fmt.Errorf("fail read config: %v", err).Error())
		os.Exit(1)
	}

	return instance
}

func (c *Config) SetMinioPassword(psw string) {
	c.Minio.RootPsw = psw
}
