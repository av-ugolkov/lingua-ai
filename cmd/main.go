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

	var minioPsw string
	flag.StringVar(&minioPsw, "minio_psw", runtime.EmptyString, "password for minio")

	flag.Parse()

	if minioPsw == runtime.EmptyString {
		panic("empty minio_psw")
	}

	cfg := config.Init(configPath)
	cfg.SetMinioPassword(minioPsw)

	app.ServerStart(cfg)
}
