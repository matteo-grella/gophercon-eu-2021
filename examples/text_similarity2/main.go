package main

import (
	"bufio"
	"fmt"
	"github.com/nlpodyssey/spago/pkg/mat32"
	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
	"io"
	"log"
	"os"
	"sort"
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

	sentences := getSentencesFromFile(os.Args[2])
	vectors := make([][]float32, len(sentences))
	for i, sentence := range sentences {
		vectors[i] = vectorize(sentence) // can be concurrent
	}

	fn := func(text string) {
		hits := rankBySimilarity(vectors, vectorize(text))
		for i, item := range limit(hits, 5) {
			fmt.Printf("%d. %s [%s]\n", i, sentences[item.id], colorize(item.score))
		}
		fmt.Println()
	}
	forEachInput(os.Stdin, fn)
}

func limit(lst []idScorePair, max int) []idScorePair {
	return lst[:min(max, len(lst))]
}

type idScorePair struct {
	id    int
	score float32
}

// Note: a fast approximate nearest neighbor search would be perfect for the job.
func rankBySimilarity(vectors [][]float32, query []float32) []idScorePair {
	hits := make([]idScorePair, 0, len(vectors))
	for id, vector := range vectors {
		hits = append(hits, idScorePair{
			id:    id,
			score: dotProduct(query, vector),
		})
	}
	sort.Slice(hits, func(i, j int) bool {
		return hits[j].score < hits[i].score
	})
	return hits
}

func dotProduct(v1, v2 []float32) float32 {
	var s float32 = 0
	_ = v2[len(v1)-1] // avoid bounds check
	for i, a := range v1 {
		s += a * v2[i]
	}
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func getSentencesFromFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
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
