package config

import (
	"fmt"
	"time"
)

type Service struct {
	Port           uint16   `yaml:"port" env-default:"5001"`
	AllowedOrigins []string `yaml:"allowed_origins" env-default:"http://localhost:5000"`
}

type DbSql struct {
	Name              string `yaml:"name"`
	User              string `yaml:"user"`
	Password          string `yaml:"-"`
	Host              string `yaml:"host"`
	Port              uint16 `yaml:"port"`
	MaxConns          uint16 `yaml:"max_conns"`
	MinConns          uint16 `yaml:"min_conns"`
	MaxConnLifetime   uint32 `yaml:"max_conn_life_time"`
	MaxConnIdleTime   uint32 `yaml:"max_conn_idle_time"`
	HealthCheckPeriod uint32 `yaml:"health_check_period"`
	ConnectTimeout    uint32 `yaml:"connect_timeout"`
}

type (
	Tts struct {
		Debug              uint8            `yaml:"debug"`
		NumThreads         uint8            `yaml:"num-threads"`
		Sid                uint             `yaml:"sid"`
		TtsMaxNumSentences uint             `yaml:"tts-max-num-sentences"`
		Provider           string           `yaml:"provider"`
		TtsRuleFsts        string           `yaml:"tts-rule-fsts"`
		TtsRuleFars        string           `yaml:"tts-rule-fars"`
		Models             map[string]Model `yaml:"models"`
		Timeout            time.Duration    `yaml:"timeout"`
	}

	Model struct {
		VitsNoiseScale  float32 `yaml:"vits-noise-scale"`
		VitsNoiseScaleW float32 `yaml:"vits-noise-scale-w"`
		VitsLengthScale float32 `yaml:"vits-length-scale"`
		VitsModel       string  `yaml:"vits-model"`
		VitsLexicon     string  `yaml:"vits-lexicon"`
		VitsTokens      string  `yaml:"vits-tokens"`
		VitsDataDir     string  `yaml:"vits-data-dir"`
	}
)

type Minio struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	RootUser string `yaml:"root_user"`
	RootPsw  string `yaml:"-"`
}

func (m *Minio) Addr() string {
	return fmt.Sprintf("%s:%s", m.Host, m.Port)
}
