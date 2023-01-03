package main

import (
	"context"
	"log"

	pb "github.com/guergabo/pluggable-component/protos/dapr/proto/components/v1"
)

func doGet(c pb.StateStoreClient) {
	log.Printf("doGet function was invoked")
	res, err := c.Get(context.Background(), &pb.GetRequest{
		Key: "Gabriel",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Result: %s", res.Data)
}
