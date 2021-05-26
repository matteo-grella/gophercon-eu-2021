package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nlpodyssey/spago/pkg/nlp/charlm"
)

func main() {
	model, err := charlm.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	generator := charlm.NewGenerator(model,
		charlm.GeneratorConfig{
			MaxCharacters: 300, StopAtEOS: true, Temperature: 0.4,
		},
	)

	fn := func(text string) {
		result, _ := generator.GenerateText(text)
		fmt.Printf("%s\n", result)
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
