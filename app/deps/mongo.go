package deps

import (
	lib "go_scaffold/library/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoDB *mongo.Client
)

func InitMongo() {

	MongoDB = lib.New(&AppConfig.Mongo)

}
