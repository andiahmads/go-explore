package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var mgo *mongo.Client
var err error

// collections

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoConn := "mongodb://root:endi@localhost:27017/"
	clientOptions := options.Client().ApplyURI(mongoConn).SetWriteConcern(writeconcern.New(writeconcern.W(1), writeconcern.J(false)))
	// clientOptions := options.Client().ApplyURI(os.Getenv("MONGO")).SetWriteConcern(writeconcern.New(writeconcern.W(1), writeconcern.J(false)))
	mgo, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
		panic(err)
	}
	err = mgo.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
		panic(err)
	}
	// fmt.Println("mongo connect", mgo)
	mgo.Database("blogdb").Collection("blog")

}

func GetDB() *mongo.Client {
	return mgo
}
