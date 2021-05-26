package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
)

func main() {
	model, err := bert.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	fn := func(text string) {
		result := model.PredictMLM(text + " ")
		for _, token := range result {
			text = strings.Replace(text,
				"[MASK]", green(token.Text), 1)
		}
		fmt.Printf("%s\n\n", text)
	}
	forEachInput(os.Stdin, fn)
}

func green(str string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", str)
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
