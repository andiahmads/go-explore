package service

import (
	"context"
	pb "grpc/blog/proto"
	"log"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("===readBlog was Invoked======")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was Read: %v\n", res)
	return res

}
