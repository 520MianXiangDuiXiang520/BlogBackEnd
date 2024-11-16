package mgo

import (
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"JuneBlog/internal/utils"
	"JuneBlog/patch/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mgo) TagC() *mongo.Collection {
	return m.Db().Collection(CollTag)
}

func (m *Mgo) GetTags(ctx context.Context, opts ...opt.Opt) ([]*module.Tag, error) {
	cc := opt.NewAndApplyCtx(opts...)
	fOpts := options.Find()
	if cc.Page > 0 && cc.PageSize > 0 {
		fOpts.SetSkip(cc.Page * cc.PageSize)
		fOpts.SetLimit(cc.PageSize)
	}
	if cc.OrderBy != "" {
		fOpts.SetSort(bson.M{cc.OrderBy: 1})
	}
	fOpts.SetProjection(bson.M{"_id": 0})
	cur, err := m.TagC().Find(ctx, cc.ApplyFilter(), fOpts)
	if err != nil {
		return nil, HandleError(err)
	}
	result := make([]*module.Tag, 0)
	err = cur.All(ctx, &result)
	return result, HandleError(err)
}

func (m *Mgo) DelTag(ctx context.Context, id int64, opts ...opt.Opt) error {
	filter := bson.M{"id": id}
	_, err := m.TagC().DeleteOne(ctx, filter)
	return HandleError(err)
}

func (m *Mgo) NewTag(ctx context.Context, tag *module.Tag, opts ...opt.Opt) error {
	id, err := m.NextId(ctx, CollTag)
	if err != nil {
		logger.Error("Tag: gen id err", err)
		return HandleError(err)
	}
	tag.Id = id
	tag.CreateTs = utils.NowTs()
	_, err = m.TagC().InsertOne(ctx, tag)
	return HandleError(err)
}
