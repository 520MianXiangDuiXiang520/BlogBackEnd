package mgo

import (
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"JuneBlog/patch/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mgo) TokenC() *mongo.Collection {
	return m.Db().Collection(CollToken)
}

func (m *Mgo) SetToken(ctx context.Context, token string, opts ...opt.Opt) error {
	filter := bson.M{}
	_, err := m.TokenC().DeleteMany(ctx, filter)
	if err != nil {
		logger.Error("Db: delete old token fail", "err", err)
		return HandleError(err)
	}
	_, err = m.TokenC().InsertOne(ctx, module.TokenC{Token: token})
	if err != nil {
		logger.Error("Db: insert new token fail", "err", err)
		return HandleError(err)
	}
	return nil
}

func (m *Mgo) GetToken(ctx context.Context, opts ...opt.Opt) (string, error) {
	cur, err := m.TokenC().Find(ctx, bson.M{}, options.Find().SetLimit(1))
	if err != nil {
		logger.Error("Db: get token fail", "err", err)
		return "", HandleError(err)
	}
	result := make([]*module.TokenC, 0)
	err = cur.All(ctx, &result)
	if err != nil {
		logger.Error("Db: unmarshal token fail", "err", err)
		return "", HandleError(err)
	}
	if len(result) == 0 {
		return "", nil
	}

	return result[0].Token, nil
}

func (m *Mgo) DelToken(ctx context.Context, opts ...opt.Opt) error {
	filter := bson.M{}
	_, err := m.TokenC().DeleteMany(ctx, filter)
	if err != nil {
		logger.Error("Db: delete old token fail", "err", err)
		return HandleError(err)
	}
	return nil
}
