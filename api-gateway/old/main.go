package main

import (
	"context"
	"log"

	pbExample "github.com/u-shylianok/ad-service/svc-ads/client/example"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	exampleClient := pbExample.NewExampleServiceClient(nil)

	result, err := exampleClient.ExampleFunc(context.Background(), &pbExample.ExampleRequest{Value: 24})
	log.Println(result)
	log.Println(err)
}
