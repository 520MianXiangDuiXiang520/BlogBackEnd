package mgo

type CollectionName string

const (
	CollInternalIdGenerate = "internal_mongo_utils_id_generator"
	CollArticle            = "article"
	CollTag                = "tag"
	CollToken              = "token"
)

var c2m = make(map[string]struct{})

func init() {
	c2m[CollArticle] = struct{}{}
	c2m[CollTag] = struct{}{}
	c2m[CollInternalIdGenerate] = struct{}{}
	c2m[CollToken] = struct{}{}
}
