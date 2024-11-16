package mgo

import (
	"JuneBlog/internal/common"
	"JuneBlog/patch/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func Inc(key string, n int64) bson.M {
	return bson.M{"$inc": bson.M{key: n}}
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}
	logger.Error("handle err", err)
	return common.SrvErrorDb
}
