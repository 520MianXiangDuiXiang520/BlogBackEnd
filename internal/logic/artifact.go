package logic

import (
	"JuneBlog/internal/common"
	"JuneBlog/internal/config"
	"JuneBlog/internal/db"
	"JuneBlog/internal/db/module"
	"JuneBlog/internal/db/opt"
	"JuneBlog/internal/message"
	"JuneBlog/internal/utils"
	"JuneBlog/patch/ginx"
	"JuneBlog/patch/logger"
	"context"
	"strings"
	"unicode/utf8"
)

func ArticleListReq(ctx context.Context, page, pageSize int, _ message.None) (resp *message.ArticleListResp, err error) {
	ginCtx := ginx.GinCtx(ctx)
	tag := int32(ginx.QueryDefaultInt(ginCtx, "tag", 0))
	countOpt := []opt.Opt{}
	opts := []opt.Opt{
		opt.WithPage(int64(page)),
		opt.WithPageSize(int64(pageSize)),
		opt.WithOrderBy("id", true),
	}
	if tag > 0 {
		tagOpt := opt.WithFilter(opt.In("tag_ids", []int32{tag}))
		opts = append(opts, tagOpt)
		countOpt = append(countOpt, tagOpt)
	}

	articles, err := db.Db().FindSomeArticleInfo(ctx, opts...)
	if err != nil {
		logger.Error("Article: FindSomeArticleInfo handle fail", "err", err)
		return nil, err
	}
	total, err := db.Db().FindArticleCount(ctx, countOpt...)
	if err != nil {
		logger.Error("Article: FindArticleCount handle fail", "err", err)
		return nil, err
	}
	resp = &message.ArticleListResp{
		RespHeader: message.RespOK,
		Articles:   articles,
		TotalSize:  total,
	}
	return resp, nil
}

func splitAbstract(text string) string {
	abstractList := strings.Split(text, common.AbstractSplitStr)
	sp := config.G.GetIntWithDefault(config.CommonKeyAbstractLen, 200)
	// 没有显示定义摘要，提取文字前部分内容作为摘要
	if len(abstractList) < 2 {
		if utf8.RuneCountInString(text) > sp {
			str := string([]rune(text)[:sp]) + "..."
			str = utils.RemoveTitle(str)
			return strings.Replace(str, "\n", "", -1)
		}
		// 文章很短的情况
		str := utils.RemoveTitle(text)
		return strings.Replace(str, "\n", "", -1)
	}

	r := utils.RemoveTitle(abstractList[0])
	r = strings.Replace(r, "\n", "", len(r))
	if utf8.RuneCountInString(r) > sp {
		return string([]rune(r)[:sp-3]) + "..."
	}
	return r
}

func NewArtifactReq(ctx context.Context, req message.NewArticleReq) (resp *message.RespHeader, err error) {
	abstract := req.Abstract
	if abstract == "" {
		abstract = splitAbstract(req.Text)
	}

	now := utils.NowTs()
	article := &module.Article{
		Header: module.ArticleHeader{
			Name:     req.Title,
			CreateTs: now,
			UpdateTs: now,
			Abstract: abstract,
			TagIds:   utils.SliceConversion[int32, int64](req.Tags),
		},
		Text: req.Text,
	}
	err = db.Db().NewArticle(ctx, article)
	if err != nil {
		logger.Error("Article: NewArticle fail", "err", err)
		return nil, err
	}
	return &message.RespOK, nil
}

func ArticleDetailReq(ctx context.Context, id int, _ message.None) (resp *message.ArticleDetailResp, err error) {
	article, err := db.Db().FindOneArticleDetail(ctx, int64(id))
	if err != nil {
		logger.Error("Article: FindOneArticle fail", "err", err)
		return nil, err
	}
	resp = &message.ArticleDetailResp{
		RespHeader: message.RespOK,
		Title:      article.Header.Name,
		CreateTs:   article.Header.CreateTs,
		UpdateTs:   article.Header.UpdateTs,
		Text:       article.Text,
		Tags:       utils.SliceConversion[int64, int32](article.Header.TagIds),
	}
	return resp, nil
}

func ArticleUpdateReq(
	ctx context.Context, id int, req message.NewArticleReq,
) (resp *message.RespHeader, err error) {
	article, err := db.Db().FindOneArticleDetail(ctx, int64(id))
	if err != nil {
		logger.Error("Article: FindOneArticle fail", "err", err)
		return nil, err
	}

	article.Text = req.Text
	article.Header.Name = req.Title
	article.Header.UpdateTs = utils.NowTs()
	article.Header.TagIds = utils.SliceConversion[int32, int64](req.Tags)
	abstract := req.Abstract
	if req.Abstract == "" {
		abstract = splitAbstract(req.Text)
	}
	article.Header.Abstract = abstract

	err = db.Db().UpdateArticle(ctx, int64(id), article)
	if err != nil {
		logger.Error("Article: UpdateArticle fail", "err", err)
		return nil, err
	}
	return &message.RespOK, nil
}

func ArticleDeleteReq(
	ctx context.Context, id int, req message.None,
) (resp *message.RespHeader, err error) {
	err = db.Db().DelArticle(ctx, int64(id))
	if err != nil {
		logger.Error("Article: DeleteArticle fail", "err", err, "id", id)
		return nil, err
	}
	return &message.RespOK, nil
}
