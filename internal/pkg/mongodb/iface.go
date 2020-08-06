package mongodb

import "go.mongodb.org/mongo-driver/mongo/options"

type Iface interface {
	FindOne(filter interface{}, result interface{}, opts ...*options.FindOneOptions) (err error)
	Find(filter interface{}, result interface{}, opts ...*options.FindOptions) (err error)
	Insert(data []interface{}, opts ...*options.InsertManyOptions) (err error)
	InsertOne(data interface{}, opts ...*options.InsertOneOptions) (err error)
	Count(filter interface{}, opts ...*options.CountOptions) (count int64, err error)
}
