package main

import (
	"bufio"
	"fmt"
	"github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	model, err := bert.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	vectorize := func(text string) []float32 {
		vector, err := model.Vectorize(text, bert.ClsToken)
		if err != nil {
			log.Fatal(err)
		}
		return vector.(*mat32.Dense).Normalize2().Data()
	}

	fn := func(text string) {
		text1, text2 := splitByPipe(text)
		score := dotProduct(vectorize(text1), vectorize(text2))
		fmt.Printf("Similarity: %s\n", colorize(score))
	}
	forEachInput(os.Stdin, fn)
}

func dotProduct(v1, v2 []float32) float32 {
	var s float32 = 0
	_ = v2[len(v1)-1] // avoid bounds check
	for i, a := range v1 {
		s += a * v2[i]
	}
	return s
}

func colorize(score float32) string {
	switch {
	case score >= 0.7:
		return fmt.Sprintf("\033[1;32m%.2f\033[0m", score) // green
	case score >= 0.55:
		return fmt.Sprintf("\033[1;33m%.2f\033[0m", score) // yellow
	default:
		return fmt.Sprintf("\033[1;31m%.2f\033[0m", score) // red
	}
}

func splitByPipe(text string) (text1, text2 string) {
	spl := strings.Split(text, "|")
	if len(spl) < 2 {
		log.Fatal(fmt.Errorf("invalid text without pipe"))
	}
	return spl[0], spl[1]
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
