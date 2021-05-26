# Matteo Grella @ GopherCon Europe 2021

At [GopherCon Europe 2021](https://gophercon.eu/), Matteo Grella discussed the [spaGO](https://github.com/nlpodyssey/spago/) package, the first and only pure Go library that focuses on cutting-edge neural technologies for Natural Language Processing (NLP).

He gave some straightforward examples of how to use the library in his [talk](https://gophercon.eu/schedule/#session-25).

This repository contains those examples (and a few more) for integrating ML/NLP features and functionality into your Go programs.

# Getting Started

Here is a list of available demos (each one has its own README):

- [Expression Graph](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/expression_graph)
- [Linear Regression](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/linear_regression)
- [Named Entities Recognition](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/named_entities_recognition)
- [Named Entities Recognition (gRPC)](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/ner_grpc)
- [Character Language Model](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/character_language_model)
- [Masked Language Model](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/masked_language_model)
- [Question-Answering](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/question_answering)
- [Zero-Shot Text Classification](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/zero_shot_text_classification)
- [Cross-Lingual Text Similarity](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/text_similarity)
- [Cross-Lingual Text Similarity (2)](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/text_similarity2)
- [Machine Translation](https://github.com/matteo-grella/gophercon-eu-2021/tree/main/examples/machine_translation)

You **must** download the pre-trained neural models before you run the demos.

For this purpose, you can use the following scripts, which handle the downloading and conversion of Flair and Hugging Face models.

> Each model weighs about 2 Gb, so make sure you have enough space on your hard disk!

### Flair models

Run the script:

```console
./download_flair.sh models
```

### Hugging Face models

This script uses the Hugging Face Model Importer tool provided by spaGO.

We provide you with a `hf-importer` binary compiled for Linux AMD64 in this repository.

Here's the command used to compile it from the spaGO sources:

```console
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-extldflags=-static" -o hf-importer cmd/huggingfaceimporter/main.go
```

Alternatively, you can build it by yourself following the [instructions](https://github.com/nlpodyssey/spago/tree/main/cmd/huggingfaceimporter) in the spaGO repo.

Then, run the script:

```console
./download_hf.sh models
```

Enjoy ;)
