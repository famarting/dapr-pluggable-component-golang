package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/guergabo/pluggable-component/protos/dapr/proto/components/v1"
)

var socket string = "unix:////tmp/dapr-components-sockets/my-component.sock"

func main() {
	// insecure connection
	opts := []grpc.DialOption{}
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, creds)

	// create connection with gRPC function
	conn, err := grpc.Dial(socket, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// executable for the client
	c := pb.NewStateStoreClient(conn)
	doGet(c)
}
