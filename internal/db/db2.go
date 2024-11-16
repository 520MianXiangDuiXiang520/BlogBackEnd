package db

import (
	"JuneBlog/internal/db/logic/storage/mgo"
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"context"
)

var (
	_ IDbOperation = (*DbLogic)(nil)
)

type DbLogic struct {
	c IDbOperation
	s IDbOperation
}

func NewDbLogic() (*DbLogic, error) {
	s, err := mgo.NewMgo()
	if err != nil {
		return nil, err
	}
	d := &DbLogic{}
	d.s = s
	d.c = nil
	return d, nil
}

func (d *DbLogic) Cache() IDbOperation {
	return d.c
}

func (d *DbLogic) Storage() IDbOperation {
	return d.s
}

//////////////////////// id generation ////////////////////////
/////

func (d *DbLogic) NextId(ctx context.Context, coll string) (int64, error) {
	return d.s.NextId(ctx, coll)
}

func (d *DbLogic) ResetId(ctx context.Context, coll string) error {
	return d.s.ResetId(ctx, coll)
}

func (d *DbLogic) SetId(ctx context.Context, coll string, id int64) error {
	return d.s.SetId(ctx, coll, id)
}

///////////////////////////// tag /////////////////////////////
/////

func (d *DbLogic) GetTags(ctx context.Context, opts ...opt.Opt) ([]*module.Tag, error) {
	return d.Storage().GetTags(ctx, opts...)
}

func (d *DbLogic) DelTag(ctx context.Context, id int64, opts ...opt.Opt) error {
	return d.Storage().DelTag(ctx, id, opts...)
}

func (d *DbLogic) NewTag(ctx context.Context, tag *module.Tag, opts ...opt.Opt) error {
	return d.Storage().NewTag(ctx, tag, opts...)
}

/////////////////////////// article ///////////////////////////
/////

func (d *DbLogic) FindArticleCount(ctx context.Context, opts ...opt.Opt) (int, error) {
	//if n, err := d.Cache().FindArticleCount(nil, opts...); err == nil {
	//	return n, nil
	//}
	// cache miss
	return d.Storage().FindArticleCount(ctx, opts...)
}

func (d *DbLogic) FindSomeArticleInfo(ctx context.Context, opts ...opt.Opt) ([]*module.ArticleHeader, error) {
	return d.Storage().FindSomeArticleInfo(ctx, opts...)
}

func (d *DbLogic) FindOneArticleInfo(ctx context.Context, id int64, opts ...opt.Opt) (*module.ArticleHeader, error) {
	return d.Storage().FindOneArticleInfo(ctx, id, opts...)
}

func (d *DbLogic) FindOneArticleDetail(ctx context.Context, id int64, opts ...opt.Opt) (*module.Article, error) {
	return d.Storage().FindOneArticleDetail(ctx, id, opts...)
}

func (d *DbLogic) HasArticle(ctx context.Context, id int64, opts ...opt.Opt) bool {
	return d.Storage().HasArticle(ctx, id, opts...)
}

func (d *DbLogic) NewArticle(ctx context.Context, article *module.Article, opts ...opt.Opt) error {
	return d.Storage().NewArticle(ctx, article, opts...)
}

func (d *DbLogic) UpdateArticle(ctx context.Context, id int64, article *module.Article, opts ...opt.Opt) error {
	return d.Storage().UpdateArticle(ctx, id, article, opts...)
}

func (d *DbLogic) DelArticle(ctx context.Context, id int64, opts ...opt.Opt) error {
	return d.Storage().DelArticle(ctx, id, opts...)
}

/////////////////////////// token ///////////////////////////
/////

func (d *DbLogic) SetToken(ctx context.Context, token string, opts ...opt.Opt) error {
	return d.Storage().SetToken(ctx, token, opts...)
}

func (d *DbLogic) GetToken(ctx context.Context, opts ...opt.Opt) (string, error) {
	return d.Storage().GetToken(ctx, opts...)
}

func (d *DbLogic) DelToken(ctx context.Context, opts ...opt.Opt) error {
	return d.Storage().DelToken(ctx, opts...)
}
