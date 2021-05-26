package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bart/tasks/zsc"
)

func main() {
	model, err := zsc.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	// arbitrary list of topics
	classes := []string{"positive", "negative"}

	fn := func(text string) {
		result, err := model.Classify(text, "", classes, false)
		if err != nil {
			log.Fatal(err)
		}
		for i, item := range result.Distribution {
			fmt.Printf("%d. %s [%.2f]\n", i, item.Class, item.Confidence)
		}
		fmt.Println()
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
