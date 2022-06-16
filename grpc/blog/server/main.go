package main

import (
	pb "grpc/blog/proto"
	"grpc/blog/server/config"
	"grpc/blog/server/service"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)



var addr = "0.0.0.0:50011"
var client *mongo.Client

func main() {

	var server service.Server
	//mongo conn

	config.Init()

	//create connection GRPC
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}


