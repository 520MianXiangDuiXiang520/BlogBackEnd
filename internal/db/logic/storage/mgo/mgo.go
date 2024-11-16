package mgo

import (
	"JuneBlog/internal/config"
	"JuneBlog/patch/logger"
	"context"
	"github.com/520MianXiangDuiXiang520/duckCfg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"slices"
	"time"
)

type Mgo struct {
	cli *mongo.Client
	db  *mongo.Database
}

func NewMgo() (*Mgo, error) {
	m := &Mgo{}
	err := m.init()
	if err != nil {
		return nil, err
	}
	return m, nil
}

type MongoCfg struct {
	URI           string   `json:"uri"`
	Database      string   `json:"database"`
	Compressors   []string `json:"compressors"`
	MaxConnecting uint64   `json:"max_connecting"`
	MaxPoolSize   uint64   `json:"max_pool_size"`
	MinPoolSize   uint64   `json:"min_pool_size"`
	ConnTimeout   int64    `json:"conn_timeout"` // s
}

func (m *Mgo) newConn() error {
	cfg := config.G.Db.Mgo
	opts := options.Client()
	opts.ApplyURI(cfg.Uri)
	opts.SetCompressors(cfg.Compressors)
	opts.SetMaxConnecting(cfg.MaxConnecting)
	opts.SetMaxPoolSize(cfg.MaxPoolSize)
	opts.SetMinPoolSize(cfg.MinPoolSize)
	ctx, done := context.WithTimeout(context.Background(),
		time.Duration(cfg.ConnTimeout)*time.Second)
	defer done()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Error("connect mongo fail",
			"key", cfg.Uri, err)
		return err
	}
	m.cli = client
	return nil
}

func (m *Mgo) Cli() *mongo.Client {
	return m.cli
}

func (m *Mgo) Db() *mongo.Database {
	if m.db != nil {
		return m.db
	}
	dbName := duckcfg.Cfg().GetStringDefault("db.mongo.database", "juneBlog")
	//opts := options.Database().SetReadConcern(readconcern.ReadConcern{})
	m.db = m.Cli().Database(dbName)
	return m.db
}

func (m *Mgo) init() error {
	err := m.newConn()
	if err != nil {
		return err
	}
	ctx, done := context.WithTimeout(context.Background(), time.Second)
	defer done()
	list, err := m.Db().ListCollectionNames(ctx, bson.M{},
		options.ListCollections().SetNameOnly(true))
	if err != nil {
		logger.Error("fail to list collection", err)
		return err
	}
	for collectionName, _ := range c2m {
		if _, ok := slices.BinarySearch(list, collectionName); !ok {
			err = m.Db().CreateCollection(context.Background(), collectionName)
			if err != nil {
				logger.Error("fail to create collection", "name", collectionName,
					"err", err)
				return err
			}
		}
	}
	return nil
}
