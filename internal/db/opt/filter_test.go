package opt

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestFilter(t *testing.T) {
	pageFilter := Eq("page", 1)
	pm := pageFilter.ToBsonM()
	assert.Equal(t, pm["page"], 1)
	assert.Equal(t, len(pm), 1)

	pageSizeFilter := Eq("pageSize", 10)
	pm = pageSizeFilter.ToBsonM()
	assert.Equal(t, pm["pageSize"], 10)
	assert.Equal(t, len(pm), 1)

	andFilter := And(pageFilter, pageSizeFilter)
	pm = andFilter.ToBsonM()
	assert.Equal(t, pm["$and"], bson.A{pageFilter.ToBsonM(), pageSizeFilter.ToBsonM()})
	assert.Equal(t, len(pm), 1)

	inFilter := In("tag", []int{1, 2, 3})
	pm = inFilter.ToBsonM()
	assert.Equal(t, pm["tag"], bson.M{"$in": []int{1, 2, 3}})
	assert.Equal(t, len(pm), 1)

	orFilter := Or(andFilter, inFilter)
	pm = orFilter.ToBsonM()
	assert.Equal(t, pm["$or"], bson.A{andFilter.ToBsonM(), inFilter.ToBsonM()})
	assert.Equal(t, len(pm), 1)
}

func TestFilter2(t *testing.T) {
	f := Or(Eq("page", 1).And(Eq("pageSize", 10)), In("tag", []int{1, 2, 3}))
	pm := f.ToBsonM()
	assert.Equal(t, pm,
		bson.M{"$or": bson.A{
			bson.M{"$and": bson.A{
				bson.M{"page": 1},
				bson.M{"pageSize": 10},
			}},
			bson.M{"tag": bson.M{"$in": []int{1, 2, 3}},
			},
		}})
}
