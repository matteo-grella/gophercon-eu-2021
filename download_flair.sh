#!/bin/sh

set -e

download_model() {
  name=$1
  url=$2

  if [ -d "$name" ]; then
    echo "Model $name already exists."
    return
  fi

  echo "Downloading model $name..."
  wget --progress=bar -O "$name.tar.gz" "$url"

  echo "Extracting model $name..."
  tar zxf "$name.tar.gz" "$name"

  rm "$name.tar.gz"
}

models_path=${1:-.}
mkdir -p "$models_path"
cd "$models_path"

d='https://dl.dropboxusercontent.com/s'

# NER English (ONTONOTES)
# download_model 'goflair-en-ner-ontonotes-fast-v0.4' "$d/a77mfbr1mvzqzcr/goflair-en-ner-ontonotes-fast-v0.4.tar.gz?dl=0"

# NER English (CONLL 2003)
# download_model 'goflair-en-ner-conll03-v0.4' "$d/uf1jihxxb5lsyvy/goflair-en-ner-conll03-v0.4.tar.gz?dl=0"

# NER English (CONLL 2003)
# download_model 'goflair-en-ner-fast-conll03-v0.4' "$d/pu53gqxlpuzmmwr/goflair-en-ner-fast-conll03-v0.4.tar.gz?dl=0"

# NER French (WIKI)
# download_model 'goflair-fr-ner-wikiner-0.4' "$d/588byt40sc5v1vo/goflair-fr-ner-wikiner-0.4.tar.gz?dl=0"

# NER MULTILINGUAL (CONLL 2003)
# download_model 'goflair-ner-multi-fast' "$d/z1nr57a8zh4qwml/goflair-ner-multi-fast.tar.gz?dl=0"

# Chunk English (CONLL 2000)
# download_model 'goflair-en-chunk-conll2000-fast-v0.4' "$d/m3sraek9iy663gp/goflair-en-chunk-conll2000-fast-v0.4.tar.gz?dl=0"

# Character Language Model - English (WIKI)
download_model 'goflair-news-forward-0.4.1' "$d/sq9kzdnd2etdmv0/goflair-news-forward-0.4.1.tar.gz?dl=0"

echo 'Done.'
