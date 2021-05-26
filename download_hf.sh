#!/usr/bin/env bash
set -e
models_path=${1:-.}
mkdir -p "$models_path"
model_names=(
  'bert-base-cased' # masked language model
  'deepset/bert-base-cased-squad2' # question answering
  'valhalla/distilbart-mnli-12-3' # natural language inference
  'pvl/labse_bert' # Google Language-agnostic BERT Sentence Embedding
  'Helsinki-NLP/opus-mt-it-en' # machine translation (italian to english)
)
for model_name in "${model_names[@]}"; do
  ./hf-importer --repo "$models_path" --model "$model_name"
done
