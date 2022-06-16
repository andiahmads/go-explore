package main

import (
	"log"

	"grpc/blog/client/service"
	pb "grpc/blog/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50011"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := service.CreateBlog(c)

	service.ReadBlog(c, id)

	// readBlog(c, id)
	// readBlog(c, "aNonExistingID")
	// updateBlog(c, id)
	// listBlog(c)
	// deleteBlog(c, id)
}

// func createBlog(c pb.BlogServiceClient) string {
// 	log.Println("---createBlog was invoked---")

// 	blog := &pb.Blog{
// 		AuthorId: "andi ahmad",
// 		Title:    "My First Blog",
// 		Content:  "Content of the first blog",
// 	}

// 	res, err := c.CreateBlog(context.Background(), blog)

// 	if err != nil {
// 		log.Fatalf("Unexpected error: %v\n", err)
// 	}

// 	log.Printf("Blog has been created: %v\n", res)
// 	return res.Id
// }
