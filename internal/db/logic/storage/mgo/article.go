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

func (m *Mgo) ArticleC() *mongo.Collection {
	return m.Db().Collection(CollArticle)
}

// onlyHeaderOpt 只读文章概要信息
var onlyHeaderOpt = bson.M{"talk_ids": 0, "text": 0}

func (m *Mgo) FindArticleCount(ctx context.Context, opts ...opt.Opt) (int, error) {
	cc := opt.NewAndApplyCtx(opts...)
	fOpts := options.Count()
	if cc.Page > 0 && cc.PageSize > 0 {
		fOpts.SetSkip(cc.Page * cc.PageSize)
		fOpts.SetLimit(cc.PageSize)
	}
	filter := cc.ApplyFilter()
	logger.Info("Mgo: FindArticleCount", "filter", filter)
	count, err := m.ArticleC().CountDocuments(ctx, filter, fOpts)
	return int(count), HandleError(err)
}

func (m *Mgo) FindSomeArticleInfo(ctx context.Context, opts ...opt.Opt) ([]*module.ArticleHeader, error) {
	cc := opt.NewAndApplyCtx(opts...)
	fOpts := options.Find()
	if cc.Page > 0 && cc.PageSize > 0 {
		fOpts.SetSkip((cc.Page - 1) * cc.PageSize)
		fOpts.SetLimit(cc.PageSize)
	}
	if cc.OrderBy != "" {
		desc := 1
		if cc.OrderDesc {
			desc = -1
		}
		fOpts.SetSort(bson.M{cc.OrderBy: desc})
	}
	fOpts.SetProjection(onlyHeaderOpt)
	filter := cc.ApplyFilter()
	logger.Info("Mgo: FindSomeArticleInfo", "filter", filter)
	cur, err := m.ArticleC().Find(ctx, filter, fOpts)
	if err != nil {
		return nil, HandleError(err)
	}
	result := make([]*module.Article, 0)
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, HandleError(err)
	}
	r := make([]*module.ArticleHeader, len(result))

	for i, article := range result {
		t := article
		//tt := t.(*module.Article)
		r[i] = &t.Header
	}
	return r, nil
}

func (m *Mgo) FindOneArticleInfo(ctx context.Context, id int64, opts ...opt.Opt) (*module.ArticleHeader, error) {
	opts = append(opts, opt.WithFilter(opt.Eq("id", id)))
	cc := opt.NewAndApplyCtx(opts...)
	filter := cc.ApplyFilter()
	fOpts := options.FindOne()
	fOpts.SetProjection(onlyHeaderOpt)
	res := m.ArticleC().FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, HandleError(res.Err())
	}
	result := &module.Article{}
	err := res.Decode(res)
	if err != nil {
		return nil, HandleError(err)
	}
	return &result.Header, nil
}

func (m *Mgo) FindOneArticleDetail(ctx context.Context, id int64, opts ...opt.Opt) (*module.Article, error) {
	opts = append(opts, opt.WithFilter(opt.Eq("id", id)))
	cc := opt.NewAndApplyCtx(opts...)
	filter := cc.ApplyFilter()
	res := m.ArticleC().FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, HandleError(res.Err())
	}

	result := &module.Article{}
	err := res.Decode(result)
	if err != nil {
		return nil, HandleError(err)
	}
	return result, nil
}

func (m *Mgo) HasArticle(ctx context.Context, id int64, opts ...opt.Opt) bool {
	opts = append(opts, opt.WithFilter(opt.Eq("id", id)))
	n, err := m.FindArticleCount(ctx, opts...)
	if err != nil {
		return false
	}
	return n > 0
}

func (m *Mgo) NewArticle(ctx context.Context, artifact *module.Article, opts ...opt.Opt) error {
	now := utils.NowTs()
	id, err := m.NextId(ctx, CollArticle)
	if err != nil {
		logger.Error("Artifact: gen id err", err)
		return HandleError(err)
	}
	artifact.Header.Id = id
	artifact.Header.CreateTs = now
	artifact.Header.UpdateTs = now
	_, err = m.ArticleC().InsertOne(ctx, artifact)
	if err != nil {
		logger.Error("Artifact: insert fail", err)
		return HandleError(err)
	}
	return nil
}

func (m *Mgo) UpdateArticle(ctx context.Context, id int64, article *module.Article, opts ...opt.Opt) error {
	filter := bson.M{"id": id}
	now := utils.NowTs()
	article.Header.UpdateTs = now
	_, err := m.ArticleC().UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: article}})
	return HandleError(err)
}

func (m *Mgo) DelArticle(ctx context.Context, id int64, opts ...opt.Opt) error {
	filter := bson.M{"id": id}
	_, err := m.ArticleC().DeleteOne(ctx, filter)
	return HandleError(err)
}
