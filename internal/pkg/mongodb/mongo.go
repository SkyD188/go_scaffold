package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoDB() *mongo.Client {
	return mongoDB
}

type mongoCollection struct {
	Timeout    time.Duration
	Collection *mongo.Collection
}

// mongo操作入口
func MongoDBCurd(database string, collection string, opts ...*options.CollectionOptions) Iface {
	dbCollection := getCollection(database, collection, opts...)
	return &mongoCollection{Timeout: 5, Collection: dbCollection}
}
func getCollection(database string, collection string, opts ...*options.CollectionOptions) *mongo.Collection {
	return getMongoDB().Database(database).Collection(collection, opts...)
}

func (m *mongoCollection) FindOne(filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error) {
	err = m.Collection.FindOne(context.Background(), filter, opts...).Decode(result)
	return
}
func (m *mongoCollection) Find(filter interface{}, result interface{}, opts ...*options.FindOptions) (err error) {
	find, err := m.Collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return err
	}
	return find.All(context.Background(), result)
}
func (m *mongoCollection) Insert(data []interface{}, opts ...*options.InsertManyOptions) (err error) {
	_, err = m.Collection.InsertMany(context.Background(), data, opts...)
	return
}
func (m *mongoCollection) InsertOne(data interface{}, opts ...*options.InsertOneOptions) (err error) {
	_, err = m.Collection.InsertOne(context.Background(), data, opts...)
	return
}
func (m *mongoCollection) Count(filter interface{}, opts ...*options.CountOptions) (count int64, err error) {
	return m.Collection.CountDocuments(context.Background(), filter, opts...)
}
