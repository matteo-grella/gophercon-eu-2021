package main

import (
	"fmt"
	"github.com/nlpodyssey/spago/pkg/nlp/sequencelabeler"
	"github.com/nlpodyssey/spago/pkg/nlp/sequencelabeler/grpcapi"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	model, err := sequencelabeler.LoadModel(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer model.Close()

	grpcServer := grpc.NewServer()
	grpcapi.RegisterSequenceLabelerServer(grpcServer, sequencelabeler.NewServer(model))

	listener, err := net.Listen("tcp", "localhost:3264")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
