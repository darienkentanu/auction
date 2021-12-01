package main

import (
	"auction/common/config"
	"auction/common/model"
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.ProductList

func init() {
	localStorage = new(model.ProductList)
	localStorage.List = make([]*model.Product, 0)
}

type AuctionServer struct{}

func (AuctionServer) AddProduct(ctx context.Context, param *model.Product) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Adding product", param.String())

	return new(empty.Empty), nil
}

func (AuctionServer) ListProduct(ctx context.Context, void *empty.Empty) (*model.ProductList, error) {
	return localStorage, nil
}

func main() {

	srv := grpc.NewServer()
	var auctionSrv AuctionServer
	model.RegisterAuctionServer(srv, auctionSrv)

	log.Println("Starting RPC server at", config.SERVICE_AUCTION_PORT)

	l, err := net.Listen("tcp", config.SERVICE_AUCTION_PORT)
	if err != nil {
		log.Fatalf("couldn't listen to %s: %v", config.SERVICE_AUCTION_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
