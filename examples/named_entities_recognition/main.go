package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nlpodyssey/spago/pkg/nlp/sequencelabeler"
)

func main() {
	model, err := sequencelabeler.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	fn := func(text string) {
		result := model.Analyze(text, true, true)
		for _, token := range result.Tokens {
			fmt.Printf("%s -> %s\n", token.Text, token.Label)
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
