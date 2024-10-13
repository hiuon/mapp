package main

import (
	"context"
	"log"
	"net"

	"hiuon/mapp/proto/userpb"

	"google.golang.org/grpc"
)

// UserServer is used to implement the UserService gRPC interface
type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

// CreateUser handles the creation of a user
func (s *UserServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := req.GetUser()
	log.Printf("Creating user: %v\n", user)

	// Normally you'd save the user to a database, here we'll just echo it back
	return &userpb.CreateUserResponse{User: user}, nil
}

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the UserServer with the gRPC server
	userpb.RegisterUserServiceServer(s, &UserServer{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
