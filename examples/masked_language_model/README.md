# Masked Language Model

This neural model performs a prediction based on a trained Masked Language Model (MLM).

In short, MLM is a fill-in-the-blank task, where the objective is to use the context words surrounding a `[MASK]` token
to try to predict what that masked word should be.

You can experiment with more `[MASK]` tokens at the same time, and the model will generate the most likely substitution
for each. Keep in mind that the more tokens are masked the less context is usable and therefore the accuracy may dro

Run this from the current directory:

```console
go run main.go ../../models/bert-base-cased/
```

Enjoy ;)

### Example

```console
> I'm so [MASK] to talk about this topic! 
I'm so excited to talk about this topic!

> The [MASK] of this neural [MASK] is impressive.
The efficiency of this neural network is impressive.

> [MASK] is a programming language
Python is a programming language

> Berlin is the capital of [MASK] .
Berlin is the capital of Germany .

> [MASK] is the capital of Germany.
Frankfurt is the capital of Germany.
```
