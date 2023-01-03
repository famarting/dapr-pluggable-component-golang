package main

import (
	"context"
	"fmt"

	pb "github.com/guergabo/pluggable-component/protos/dapr/proto/components/v1"
)

// Initializes the state store component with the given metadata.
func (s *Server) Init(ctx context.Context, in *pb.InitRequest) (*pb.InitResponse, error) {
	return &pb.InitResponse{}, nil
}

// Returns a list of implemented state store features.
func (s *Server) Features(ctx context.Context, in *pb.FeaturesRequest) (*pb.FeaturesResponse, error) {
	return &pb.FeaturesResponse{}, nil
}

// Deletes the specified key from the state store.
func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}

// Get data from the given key.
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Data: []byte(fmt.Sprintf("CREATED PLUGGABLE COMPONENT: with key %s", in.Key)),
	}, nil
}

// Sets the value of the specified key.
func (s *Server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	return &pb.SetResponse{}, nil
}

// Ping the state store. Used for liveness porpuses.
func (s *Server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

// Deletes many keys at once.
func (s *Server) BulkDelete(ctx context.Context, in *pb.BulkDeleteRequest) (*pb.BulkDeleteResponse, error) {
	return &pb.BulkDeleteResponse{}, nil
}

// Retrieves many keys at once.
func (s *Server) BulkGet(ctx context.Context, in *pb.BulkGetRequest) (*pb.BulkGetResponse, error) {
	return &pb.BulkGetResponse{}, nil
}

// Set the value of many keys at once.
func (s *Server) BulkSet(ctx context.Context, in *pb.BulkSetRequest) (*pb.BulkSetResponse, error) {
	return &pb.BulkSetResponse{}, nil
}
