#!/bin/bash

dev() {
  BRANCH="$(cut -d "/" -f2 <<< "$(git rev-parse --abbrev-ref HEAD)")"
  COMMIT="$(git $dir rev-parse HEAD)"

  minio_psw="$(cat .env | grep MINIO_PSW | cut -d "=" -f2)"
  
  BRANCH=${BRANCH} \
  COMMIT=${COMMIT} \
  MINIO_PSW=${minio_psw} \
  docker compose -p lingua-evo-dev -f deploy/docker-compose.dev.yml up --build --force-recreate    
}

release() {
  clear

  if test -d /home/lingua-logs; then
    echo "Directory [ lingua-logs ] exists."
  else
    mkdir /home/lingua-logs
  fi

  if test -d /home/lingua-dumps; then
    echo "Directory [ lingua-dumps ] exists."
  else
    mkdir /home/lingua-dumps
  fi

  BRANCH="$(cut -d "/" -f2 <<< "$(git rev-parse --abbrev-ref HEAD)")"
  echo "Do you want to deploy from [ $BRANCH ]? [y/n]"
  read ans
  if [ "$ans" = "n" ]; then
    echo "Type branch name which you want yo use"
    read BRANCH
  fi

  CURRENT_BRANCH="$(cut -d "/" -f2 <<< "$(git rev-parse --abbrev-ref HEAD)")"
  if [ "$BRANCH" != "$CURRENT_BRANCH" ]; then
    echo "$(git checkout $BRANCH)"
  fi

  echo "$(git fetch)"
  echo "$(git pull)"

  COMMIT="$(git $dir rev-parse HEAD)"

  minio_psw="$(cat .env | grep MINIO_PSW | cut -d "=" -f2)"

  BRANCH=${BRANCH} \
  COMMIT=${COMMIT} \
  MINIO_PSW=${minio_psw} \
  docker compose -p lingua-evo -f deploy/docker-compose.yml up --build --force-recreate
}

"$@"
