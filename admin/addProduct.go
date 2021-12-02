package main

import (
	"auction/common/config"
	"auction/common/model"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceAuction() model.AuctionClient {
	port := config.SERVICE_AUCTION_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewAuctionClient(conn)
}

func main() {

	product1 := model.Product{
		Id: "1", Name: "monitor", Price: 1000000,
	}

	product2 := model.Product{
		Id: "2", Name: "mouse", Price: 200000,
	}

	auction := serviceAuction()

	fmt.Println("\n", "===========> add product test")

	// add product 1
	auction.AddProduct(context.Background(), &product1)
	// add product 2
	auction.AddProduct(context.Background(), &product2)
	// view list of product
	res1, err := auction.ListProduct(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))
}
