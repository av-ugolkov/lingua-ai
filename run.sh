#!/bin/bash

run(){
    cfg="./configs/server.dev.yaml"
    pg_psw="$(cat .env | grep PG_PSW | cut -d "=" -f2)"
    minio_psw="$(cat .env | grep MINIO_PSW | cut -d "=" -f2)"

    go run ./cmd/main.go -config=${cfg} -pg_psw=${pg_psw} -minio_psw=${minio_psw}
}

"$@"
