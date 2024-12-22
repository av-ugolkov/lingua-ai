package main

import (
	"flag"

	"github.com/av-ugolkov/lingua-ai/internal/app"
	"github.com/av-ugolkov/lingua-ai/internal/config"
	"github.com/av-ugolkov/lingua-ai/runtime"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./configs/server_config.yaml", "it's name of application config")

	var pgPsw string
	flag.StringVar(&pgPsw, "pg_psw", runtime.EmptyString, "password for postgres db")

	flag.Parse()

	if pgPsw == runtime.EmptyString {
		panic("empty jwts, pg_psw or redis_psw")
	}

	cfg := config.InitConfig(configPath)
	config.SetDBPassword(pgPsw)

	app.ServerStart(cfg)
}
