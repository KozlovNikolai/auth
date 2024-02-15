package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/KozlovNikolai/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

// Create is creating new user and respons his ID
func (s *server) Create(_ context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	var temp int64
	if in.Role == desc.Role_ADMIN {
		temp = 1
	} else {
		temp = 2
	}
	fmt.Println("Create  User")
	fmt.Printf("Received name: %v\n", in.Name)
	fmt.Printf("Received email: %v\n", in.Email)
	fmt.Printf("Received password: %v\n", in.Password)
	fmt.Printf("Received password confirm: %v\n", in.PasswordConfirm)
	fmt.Printf("Received role: %v\n", in.Role)

	return &desc.CreateResponse{Id: temp}, nil
}

// Get returns data about user by his ID
func (s *server) Get(_ context.Context, in *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get User")
	var u desc.GetResponse
	start := timestamppb.Now()
	if in.Id == 3 {
		u = desc.GetResponse{
			Id:        in.Id,
			Name:      "Mike",
			Email:     "g@g.c",
			Role:      desc.Role_ADMIN,
			CreatedAt: start,
			UpdatedAt: start,
		}
	} else {
		u = desc.GetResponse{
			Id:        in.Id,
			Name:      "Pam",
			Email:     "p@p.r",
			Role:      desc.Role_USER,
			CreatedAt: start,
			UpdatedAt: start,
		}
	}

	return &u, nil
}

// Update is updating data of user
func (s *server) Update(_ context.Context, in *desc.UpdateRequest) (*empty.Empty, error) {
	fmt.Println("Update User")
	fmt.Printf("Received id: %d\n", in.Id)
	fmt.Printf("Received name: %v\n", in.Name)
	fmt.Printf("Received email: %v\n", in.Email)
	fmt.Printf("Received role: %v\n", in.Role)

	return &empty.Empty{}, nil
}

// Delete is deleting user by his ID
func (s *server) Delete(_ context.Context, in *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Println("Delete User")
	fmt.Printf("Received id: %d\n", in.Id)
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serv %v", err)
	}
}
