package mmgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountAll(ctx context.Context, colName string) (int64, error) {
	return _h.Col(colName).CountDocuments(ctx, bson.M{})
}

func Count(ctx context.Context, colName string, filter any, opts ...*options.CountOptions) (int64, error) {
	return _h.Col(colName).CountDocuments(ctx, filter, opts...)
}
