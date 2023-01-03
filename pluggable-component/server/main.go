// Pluggable State Store Example
// 1) reflection is required for service discovery
// 2) registers with Dapr via Unix Domain Sockets (using gRPC)
// 3) Test with evan-cli
//
// evans --host localhost --port 50051 --reflection repl
// sudo curl -GET --unix-socket /tmp/dapr-components-sockets/my-component.sock
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/guergabo/pluggable-component/protos/dapr/proto/components/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	socketDir     = "/tmp/dapr-components-sockets"
	componentName = "my-component"
)

type Server struct {
	pb.StateStoreServer
}

func main() {
	log.Printf("Creating pluggable component: %s\n", componentName)

	// Create directory if it does not exist
	if createDir(socketDir) {
		log.Println("Creating socket directory")
		if err := os.MkdirAll(socketDir, 0755); err != nil {
			log.Fatalf("Failed to create docker directory: %v\n", err)
		}
	}

	// Delete socket in case it already exists
	socketPath := fmt.Sprintf("%s/%s.sock", socketDir, componentName)
	if fileExists(socketPath) {
		log.Println("Removing existing socket")
		if err := os.Remove(socketPath); err != nil {
			log.Fatalf("Failed to remove existing socket: %v\n", err)
		}
	}

	// Create a Unix Domain Socket and listen for incoming connections
	log.Printf("Creating socket: %s", socketPath)
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer socket.Close()

	// Start server
	log.Printf("Starting gRPC Server")
	srv := grpc.NewServer()
	pb.RegisterStateStoreServer(srv, &Server{})
	defer srv.Stop()
	reflection.Register(srv)
	if err := srv.Serve(socket); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

func createDir(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}
