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

if [ ! -d llm/vits-piper-de_DE-thorsten-high ]; then
  mkdir llm/vits-piper-de_DE-thorsten-high
  git clone https://huggingface.co/csukuangfj/vits-piper-de_DE-thorsten-high llm/vits-piper-de_DE-thorsten-high
fi

if [ ! -d llm/vits-piper-fr_FR-siwis-medium ]; then
  mkdir llm/vits-piper-fr_FR-siwis-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-fr_FR-siwis-medium llm/vits-piper-fr_FR-siwis-medium
fi

if [ ! -d llm/vits-piper-es_ES-davefx-medium ]; then
  mkdir llm/vits-piper-es_ES-davefx-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-es_ES-davefx-medium llm/vits-piper-es_ES-davefx-medium
fi

if [ ! -d llm/vits-piper-it_IT-paola-medium ]; then
  mkdir llm/vits-piper-it_IT-paola-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-it_IT-paola-medium llm/vits-piper-it_IT-paola-medium
fi

if [ ! -d llm/vits-piper-pt_PT-tugao-medium ]; then
  mkdir llm/vits-piper-pt_PT-tugao-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-pt_PT-tugao-medium llm/vits-piper-pt_PT-tugao-medium
fi

if [ ! -d llm/vits-piper-sv_SE-nst-medium ]; then
  mkdir llm/vits-piper-sv_SE-nst-medium
  git clone https://huggingface.co/csukuangfj/vits-piper-sv_SE-nst-medium llm/vits-piper-sv_SE-nst-medium
fi

# https://huggingface.co/willwade/mms-tts-multilingual-models-onnx
# https://huggingface.co/csukuangfj?search_models=ru