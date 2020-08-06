package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type M = bson.M

var (
	mongoDB *mongo.Client
	err     error
)

func InitMongo(address string, poolSize uint64) {

	clientOptions := options.Client().ApplyURI(address)

	//连接池
	clientOptions.SetMaxPoolSize(poolSize)
	// Connect to MongoDB
	mongoDB, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = mongoDB.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mongo is Collection!!!")

}
