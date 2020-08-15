package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo
type Conf struct {
	Address     string
	MaxPoolSize uint64
}

//type mongoCollection struct {
//	Timeout    time.Duration
//	Collection *mongo.Collection
//}

func New(cfg *Conf) *mongo.Client {
	clientOptions := options.Client().ApplyURI(cfg.Address)
	//连接池
	clientOptions.SetMaxPoolSize(cfg.MaxPoolSize)
	// Connect to MongoDB
	mongoDB, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = mongoDB.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mongo is Collection!!!")
	return mongoDB
}

// mongo操作入口
//func MongoDBCurd(database string, collection string, opts ...*options.CollectionOptions) *mongoCollection {
//	dbCollection := getCollection(database, collection, opts...)
//	return &mongoCollection{Timeout: 5, Collection: dbCollection}
//}
//
//func getCollection(database string, collection string, opts ...*options.CollectionOptions) *mongo.Collection {
//	return getMongo().Database(database).Collection(collection, opts...)
//}
//
//func (m *mongoCollection) FindOne(filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error) {
//	err = m.Collection.FindOne(context.Background(), filter, opts...).Decode(result)
//	return
//}
//
//func (m *mongoCollection) Find(filter interface{}, result interface{}, opts ...*options.FindOptions) (err error) {
//	find, err := m.Collection.Find(context.Background(), filter, opts...)
//	if err != nil {
//		return err
//	}
//	return find.All(context.Background(), result)
//}
//
//func (m *mongoCollection) Insert(data []interface{}, opts ...*options.InsertManyOptions) (err error) {
//	_, err = m.Collection.InsertMany(context.Background(), data, opts...)
//	return
//}
//
//func (m *mongoCollection) InsertOne(data interface{}, opts ...*options.InsertOneOptions) (err error) {
//	_, err = m.Collection.InsertOne(context.Background(), data, opts...)
//	return
//}
//
//func (m *mongoCollection) Count(filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
//	return m.Collection.CountDocuments(context.Background(), filter, opts...)
//}
