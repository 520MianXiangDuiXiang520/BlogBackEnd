package module

type ArticleHeader struct {
	Id       int64   `bson:"id"`
	Name     string  `json:"name" bson:"name,omitempty"`
	CreateTs int64   `json:"create_ts" bson:"create_ts,omitempty"`
	UpdateTs int64   `json:"update_ts" bson:"update_ts,omitempty"`
	TagIds   []int64 `json:"tag_ids" bson:"tag_ids,omitempty"`
	Abstract string  `json:"abstract" bson:"abstract,omitempty"`
}

type Article struct {
	Header  ArticleHeader ` bson:"inline"`
	Text    string        `json:"text" bson:"text,omitempty"`
	TalkIds []int64       `json:"talk_ids" bson:"talk_ids,omitempty"`
}

func (a *Article) C() string {
	return "article"
}

func NewArticle() *Article {
	return &Article{
		Header: ArticleHeader{},
	}
}
