package test

import (
	"JuneBlog/internal/db"
	"JuneBlog/internal/db/module"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTag(t *testing.T) {
	t.Skip()
	ctx := context.Background()
	err := db.Db().NewTag(ctx, &module.Tag{Name: "ttt"})
	assert.Nil(t, err)

	tags, err := db.Db().GetTags(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tags))
	assert.Equal(t, tags[0].Name, "ttt")
}

func TestNewArtifact(t *testing.T) {
	ctx := context.Background()
	//id, err := db.Db().NextId(context.TODO(), mgo.CollArticle)
	//assert.Nil(t, err)
	//article := &module.Article{
	//	Header: module.ArticleHeader{
	//		Id:       id,
	//		Name:     "新一代实验分析引擎：驱动履约平台的数据决策",
	//		Abstract: "本文介绍了美团履约技术平台的新一代实验分析引擎，该引擎对核心实验框架进行了标准化，并融合了众多先进解决方案，有效解决小样本挑战。同时，提供了多样化的溢出效应应对策略，并针对不同业务场景提供了精准的方差和P值计算方法，以规避统计误差。希望对大家有所帮助或启发。",
	//	},
	//
	//	Text:    "Test111",
	//	TalkIds: nil,
	//}
	//err = db.Db().NewArticle(context.Background(), article)
	//assert.Nil(t, err)
	//articles, err := db.Db().FindSomeArticleInfo(context.TODO())
	//assert.Nil(t, err)
	//fmt.Printf("%#v \n", articles)
	//assert.Equal(t, 1, len(articles))
	//assert.Equal(t, "111", articles[0].Name)

	article := &module.Article{
		Header: module.ArticleHeader{
			Id:       32,
			Name:     "新一代实验分析引擎：驱动履约平台的数据决策",
			Abstract: "本文介绍了美团履约技术平台的新一代实验分析引擎，该引擎对核心实验框架进行了标准化，并融合了众多先进解决方案，有效解决小样本挑战。同时，提供了多样化的溢出效应应对策略，并针对不同业务场景提供了精准的方差和P值计算方法，以规避统计误差。希望对大家有所帮助或启发。",
		},

		Text:    "本文介绍了美团履约技术平台的新一代实验分析引擎本文介绍了美团履约技术平台的新一代实验分析引擎",
		TalkIds: nil,
	}
	err := db.Db().UpdateArticle(ctx, 32, article)
	assert.Nil(t, err)

	gotA, err := db.Db().FindOneArticleDetail(ctx, 32)
	assert.Nil(t, err)
	assert.Equal(t, gotA.Text, article.Text)
}
