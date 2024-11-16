package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

/////////////////////////// id generate ////////////////////////////

var IGPool = sync.Pool{
	New: func() any {
		return &IdGenerator{}
	},
}

type IdGenerator struct {
	Key string `json:"internal_key" bson:"internal_key"`
	Id  int64  `json:"internal_id" bson:"internal_id"`
}

func (m *Mgo) IdGenerateC() *mongo.Collection {
	return m.Db().Collection(CollInternalIdGenerate)
}

func (m *Mgo) NextId(ctx context.Context, colName string) (int64, error) {
	opt := options.FindOneAndUpdate().SetUpsert(true)
	opt.SetReturnDocument(options.After)
	res := m.IdGenerateC().FindOneAndUpdate(
		ctx, bson.M{"internal_key": colName}, Inc("internal_id", 1), opt)
	ig := IGPool.Get().(*IdGenerator)
	defer IGPool.Put(ig)
	err := res.Decode(ig)
	if err != nil {
		return 0, HandleError(err)
	}
	return ig.Id + 1, nil
}

func (m *Mgo) ResetId(ctx context.Context, colName string) error {
	_, err := m.IdGenerateC().DeleteOne(ctx, bson.M{"internal_key": colName})
	return HandleError(err)
}

func (m *Mgo) SetId(ctx context.Context, colName string, id int64) error {
	opts := options.Update().SetUpsert(true)
	_, err := m.IdGenerateC().UpdateOne(ctx,
		bson.M{"internal_key": colName},
		bson.M{"internal_id": id}, opts)
	return HandleError(err)
}
