package server

import (
	"context"
	"log"

	protos "github.com/snehalyelmati/microservices-with-go-grpc/protos/currency"
)

type Currency struct {
	l log.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l log.Logger) *Currency {
	return &Currency{l: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Println("Handle GetRate - ", "Base: ", rr.GetBase(), ", Destination: ", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
