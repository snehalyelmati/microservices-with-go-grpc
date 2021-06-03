package main

import (
	"log"
	"net"
	"os"

	protos "github.com/snehalyelmati/microservices-with-go-grpc/protos/currency"
	"github.com/snehalyelmati/microservices-with-go-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := log.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(*log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatal("Unable to listen:", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
