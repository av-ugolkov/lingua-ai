package main

import (
	"flag"

	"github.com/av-ugolkov/lingua-ai/internal/app"
	"github.com/av-ugolkov/lingua-ai/internal/config"
	"github.com/av-ugolkov/lingua-ai/runtime"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./configs/server.yaml", "it's name of application config")

	var pgPsw string
	flag.StringVar(&pgPsw, "pg_psw", runtime.EmptyString, "password for postgres db")

	var minioPsw string
	flag.StringVar(&minioPsw, "minio_psw", runtime.EmptyString, "password for minio")

	flag.Parse()

	if pgPsw == runtime.EmptyString || minioPsw == runtime.EmptyString {
		panic("empty pg_psw or minio_psw")
	}

	cfg := config.Init(configPath)
	cfg.SetDBPassword(pgPsw)
	cfg.SetMinioPassword(minioPsw)

	app.ServerStart(cfg)
}
