package main

import (
	"context"
	pb "grpc/blog/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr = "0.0.0.0:50051"

const mongUrl = "mongodb://admin:endi@localhost:27017/"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	//create conncetion MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(mongUrl))

	if err != nil {
		log.Fatalf(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	//create connection GRPC
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n ", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n ", err)

	}
}
