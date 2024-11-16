package mmgo

import "go.mongodb.org/mongo-driver/bson"

func Or(list bson.A) bson.M {
	return bson.M{"$or": list}
}

func In(key string, list bson.A) bson.M {
	return bson.M{key: bson.M{"$in": list}}
}

func Inc(key string, n int64) bson.M {
	return bson.M{"$inc": bson.M{key: n}}
}
