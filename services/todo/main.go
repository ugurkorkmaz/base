package main

import (
	"log"
	"net"

	"app/handler"

	"google.golang.org/grpc"

	todo "app/grpc/todo"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Start and listen to the Golang gRPC server

	// Specify the server address, for example, "localhost:5051"
	address := "localhost:5051"

	// Create a network connection to listen on TCP
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the TodoService server
	todo.RegisterTodoServiceServer(server, &handler.TodoServer{})

	// Register the HealthServer server
	health.RegisterHealthServer(server, &handler.HealthServer{})

	// Start the server and listen for client requests
	log.Printf("Server listening on %s\n", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
