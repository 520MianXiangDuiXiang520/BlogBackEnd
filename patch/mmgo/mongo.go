package mmgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Helper struct {
	cli     *mongo.Client
	db      *mongo.Database
	timeout time.Duration
}

var _h *Helper

func Init(cli *mongo.Client, dbName string) {
	_h = &Helper{
		cli:     cli,
		db:      cli.Database(dbName),
		timeout: time.Second,
	}
}

func (h *Helper) Db() *mongo.Database {
	return h.db
}

func (h *Helper) Col(name string) *mongo.Collection {
	return h.Db().Collection(name)
}

func (h *Helper) NewCtxTimeout() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), h.timeout)
	return ctx
}
