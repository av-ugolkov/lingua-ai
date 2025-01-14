FROM golang:1.23.4 AS builder

WORKDIR /build
COPY . .
RUN --mount=type=cache,target=/go make build

FROM debian:trixie-slim
LABEL key="Lingua AI"

ARG config_dir
ARG pg_psw
ARG minio_psw
ARG branch
ARG commit

LABEL git.branch=$branch
LABEL git.commit=$commit

WORKDIR /lingua-ai

COPY /configs/${config_dir}.yaml ./configs/server.yaml
COPY --from=builder ./build/main ./
COPY ./llm ./llm/
COPY ./lib ./lib/

EXPOSE 5100

ENV env_pg_psw=${pg_psw}
ENV env_minio_psw=${minio_psw}

ENTRYPOINT ./main -pg_psw=${env_pg_psw} -minio_psw=${env_minio_psw}
