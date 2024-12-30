#!/usr/bin/env bash

set -ex

if [ ! -d llm ]; then
  mkdir llm
fi

if [ ! -d vits-ljs ]; then
  curl -SL -o ./llm/vits-ljs.tar.bz2 https://github.com/k2-fsa/sherpa-onnx/releases/download/tts-models/vits-ljs.tar.bz2
  tar xvf ./llm/vits-ljs.tar.bz2 -C ./llm
  rm ./llm/vits-ljs.tar.bz2
fi

if [ ! -d vits-piper-en_US-lessac-medium ]; then
  curl -SL -o ./llm/vits-piper-en_US-lessac-medium.tar.bz2 https://github.com/k2-fsa/sherpa-onnx/releases/download/tts-models/vits-piper-en_US-lessac-medium.tar.bz2
  tar xvf ./llm/vits-piper-en_US-lessac-medium.tar.bz2 -C ./llm
  rm ./llm/vits-piper-en_US-lessac-medium.tar.bz2
fi

if [ ! -d vits-vctk ]; then
  curl -SL -o ./llm/vits-vctk.tar.bz2 https://github.com/k2-fsa/sherpa-onnx/releases/download/tts-models/vits-vctk.tar.bz2
  tar xvf ./llm/vits-vctk.tar.bz2 -C ./llm
  rm ./llm/vits-vctk.tar.bz2
fi

if [ ! -d vits-icefall-zh-aishell3 ]; then
  curl -SL -o ./llm/vits-icefall-zh-aishell3.tar.bz2 https://github.com/k2-fsa/sherpa-onnx/releases/download/tts-models/vits-icefall-zh-aishell3.tar.bz2
  tar xvf ./llm/vits-icefall-zh-aishell3.tar.bz2 -C ./llm
  rm ./llm/vits-icefall-zh-aishell3.tar.bz2
fi