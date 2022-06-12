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

// type BlogItem struct {
// 	ID       primitive.ObjectID `bson:"_id,omitempty"`
// 	AuthorID string             `bson:"author_id"`
// 	Title    string             `bson:"title"`
// 	Content  string             `bson:"content"`
// }

// func documentToBlog(data *BlogItem) *pb.Blog {
// 	return &pb.Blog{
// 		Id:       data.ID.Hex(),
// 		AuthorId: data.AuthorID,
// 		Title:    data.Title,
// 		Content:  data.Content,
// 	}
// }

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

// func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
// 	log.Printf("CreateBlog was invoked with %v\n", in)

// 	data := BlogItem{
// 		AuthorID: in.AuthorId,
// 		Title:    in.Title,
// 		Content:  in.Content,
// 	}

// 	res, err := collection.InsertOne(ctx, data)
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("internal server error: %v\n", err),
// 		)
// 	}

// 	oid, ok := res.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			"Cannot convert to oid",
// 		)
// 	}

// 	return &pb.BlogId{
// 		Id: oid.Hex(),
// 	}, nil

// }
