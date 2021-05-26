package main

import (
	"bufio"
	"fmt"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bart/tasks/seq2seq"
	"io"
	"log"
	"os"
)

func main() {
	model, err := seq2seq.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	fn := func(text string) {
		result, err := model.Generate(text)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n\n", result)
	}
	forEachInput(os.Stdin, fn)
}

func forEachInput(r io.Reader, fn func(text string)) {
	scanner := bufio.NewScanner(r)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			break
		}
		fn(text)
	}
}
