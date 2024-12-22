FROM golang:1.23.2-alpine AS builder

RUN --mount=type=cache,target=/var/cache/apk apk --no-cache --update --upgrade add git make

WORKDIR /build
COPY . .
RUN --mount=type=cache,target=/go make build

FROM alpine:3.20.3
LABEL key="Lingua AI"

ARG config_dir
ARG pg_psw
ARG branch
ARG commit

LABEL git.branch=$branch
LABEL git.commit=$commit

RUN --mount=type=cache,target=/var/cache/apk apk --update --upgrade add ca-certificates git bash

WORKDIR /lingua-ai

COPY /configs/${config_dir}.yaml ./configs/server.yaml
COPY --from=builder ./build/cmd/main ./

EXPOSE 5100

ENV env_pg_psw=${pg_psw}

ENTRYPOINT ./main -pg_psw=${env_pg_psw}
