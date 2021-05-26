package main

import (
	"bufio"
	"fmt"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	model, err := bert.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	paragraph := readFile(os.Args[2])

	fn := func(text string) {
		result := model.Answer(text, paragraph)
		if result != nil && result[0].Confidence < 0.5 {
			fmt.Print("Sorry, I'm not sure.\n\n")
			return
		}
		for i, answer := range result {
			fmt.Printf("%d. %s [%.2f]\n", i, answer.Text, answer.Confidence)
		}
		fmt.Println()
	}
	forEachInput(os.Stdin, fn)
}

func readFile(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
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
