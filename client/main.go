package main

import (
	"context"
	"log"
	"time"

	"hiuon/mapp/userpb"

	"google.golang.org/grpc"
)

func main() {
	// Dial the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := userpb.NewUserServiceClient(conn)

	// Prepare the user request
	user := &userpb.User{
		Id:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	req := &userpb.CreateUserRequest{User: user}

	// Call the CreateUser RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling CreateUser: %v", err)
	}

	log.Printf("User created: %v", res.User)
}
