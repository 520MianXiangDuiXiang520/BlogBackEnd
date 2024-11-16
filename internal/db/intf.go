package db

import (
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"context"
)

type IIdGenerateOperation interface {
	NextId(ctx context.Context, coll string) (int64, error)
	ResetId(ctx context.Context, coll string) error
	SetId(ctx context.Context, coll string, id int64) error
}

type IArticleDbOperation interface {
	FindArticleCount(ctx context.Context, opts ...opt.Opt) (int, error)

	FindSomeArticleInfo(ctx context.Context, opts ...opt.Opt) ([]*module.ArticleHeader, error)

	FindOneArticleInfo(ctx context.Context, id int64, opts ...opt.Opt) (*module.ArticleHeader, error)

	FindOneArticleDetail(ctx context.Context, id int64, opts ...opt.Opt) (*module.Article, error)

	HasArticle(ctx context.Context, id int64, opts ...opt.Opt) bool

	NewArticle(ctx context.Context, artifact *module.Article, opts ...opt.Opt) error

	UpdateArticle(ctx context.Context, id int64, article *module.Article, opts ...opt.Opt) error

	DelArticle(ctx context.Context, id int64, opts ...opt.Opt) error
}

type ITagDbOperation interface {
	GetTags(ctx context.Context, opts ...opt.Opt) ([]*module.Tag, error)
	DelTag(ctx context.Context, id int64, opts ...opt.Opt) error
	NewTag(ctx context.Context, tag *module.Tag, opts ...opt.Opt) error
}

type ITokenDbOperation interface {
	SetToken(ctx context.Context, token string, opts ...opt.Opt) error
	GetToken(ctx context.Context, opts ...opt.Opt) (string, error)
	DelToken(ctx context.Context, opts ...opt.Opt) error
}
