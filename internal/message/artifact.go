package message

import "JuneBlog/internal/db/module"

type Tag struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type TagListResp struct {
	RespHeader
	Tags []Tag `json:"tags"`
}

type ArticleListResp struct {
	RespHeader
	Articles  []*module.ArticleHeader
	TotalSize int
}

type NewArticleReq struct {
	Title    string  `json:"title"`
	Tags     []int32 `json:"tags"`
	Text     string  `json:"text"`
	Abstract string  `json:"abstract"`
}

type ArticleDetailResp struct {
	RespHeader
	Title    string  `json:"title"`
	CreateTs int64   `json:"create_ts"`
	UpdateTs int64   `json:"update_ts"`
	Text     string  `json:"text"`
	Tags     []int32 `json:"tags"`
}
