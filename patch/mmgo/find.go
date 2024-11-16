package mmgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne(colName string, filter, result any, opts ...*options.FindOneOptions) error {
	res := _h.Col(colName).FindOne(_h.NewCtxTimeout(), filter, opts...)
	return res.Decode(result)
}

func FindMany(colName string, filter any, result any, opts ...*options.FindOptions) error {
	cur, err := _h.Col(colName).Find(_h.NewCtxTimeout(), filter, opts...)
	if err != nil {
		return err
	}
	return cur.All(_h.NewCtxTimeout(), result)
}

func FindManyWithSort(colName string, filter any, result any, sort string,
	opts ...*options.FindOptions,
) error {
	if sort != "" {
		opts = append(opts, options.Find().SetSort(bson.D{{sort, 1}}))
	}
	return FindMany(colName, filter, result, opts...)
}

func FindManyWithLimitSkip(colName string, filter any, result any, sort string, skip, limit int64,
	opts ...*options.FindOptions,
) error {
	if sort != "" {
		opts = append(opts, options.Find().SetSort(bson.D{{sort, 1}}))
	}
	if skip > 0 {
		opts = append(opts, options.Find().SetSkip(skip))
	}
	if limit > 0 {
		opts = append(opts, options.Find().SetLimit(limit))
	}
	return FindMany(colName, filter, result, opts...)
}
