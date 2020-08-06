package mongodb

import (
	"context"
	"go_scaffold/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type M = bson.M

var (
	mongoDB *mongo.Client
	err     error
)

func InitMongo() {

	clientOptions := options.Client().ApplyURI(config.GetConf().Mongo.Address)

	//连接池
	clientOptions.SetMaxPoolSize(config.GetConf().Mongo.MaxPoolSize)
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
