package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/nlpodyssey/spago/pkg/nlp/sequencelabeler/grpcapi"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:3264", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpcapi.NewSequenceLabelerClient(conn)

	fn := func(text string) {
		result, err := client.Analyze(context.Background(),
			&grpcapi.AnalyzeRequest{
				Text:              text,
				MergeEntities:     true,
				FilterNotEntities: true,
			})
		if err != nil {
			log.Fatal(err)
		}
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
