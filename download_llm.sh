#!/usr/bin/env bash

set -ex

if [ ! -d llm ]; then
  mkdir llm
fi

if [ ! -d llm/vits-piper-en_US-lessac-high ]; then
  mkdir llm/vits-piper-en_US-lessac-high
  git clone https://huggingface.co/csukuangfj/vits-piper-en_US-lessac-high llm/vits-piper-en_US-lessac-high
fi

if [ ! -d llm/vits-piper-ru_RU-irina-medium ]; then
  mkdir llm/vits-piper-ru_RU-irina-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-ru_RU-irina-medium llm/vits-piper-ru_RU-irina-medium
fi

if [ ! -d llm/vits-piper-fi_FI-harri-medium ]; then
  mkdir llm/vits-piper-fi_FI-harri-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-fi_FI-harri-medium llm/vits-piper-fi_FI-harri-medium
fi



# https://huggingface.co/willwade/mms-tts-multilingual-models-onnx
# https://huggingface.co/csukuangfj?search_models=ru