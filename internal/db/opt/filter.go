package opt

import "go.mongodb.org/mongo-driver/bson"

type FilterTp uint16

const (
	FilterTpAnd FilterTp = iota + 1
	FilterTpOr
	FilterTpIn
	FilterTpEq
)

func And(filters ...Filter) Filter {
	return &FilterAnd{filters: filters}
}

func Or(filters ...Filter) Filter {
	return &FilterOr{filters: filters}
}

func In(key string, value any) Filter {
	return &FilterIn{key: key, value: value}
}

func Eq(key string, value any) Filter {
	return &FilterEq{key: key, value: value}
}

type Filter interface {
	GetTp() FilterTp
	GetFilters() []Filter
	ToBsonM() bson.M
	And(filter Filter) Filter
	Or(filter Filter) Filter
}

type FilterAnd struct {
	filters []Filter
}

func (f *FilterAnd) GetTp() FilterTp {
	return FilterTpAnd
}

func (f *FilterAnd) GetFilters() []Filter {
	return f.filters
}

func (f *FilterAnd) ToBsonM() bson.M {
	res := bson.A{}
	for _, f := range f.filters {
		res = append(res, f.ToBsonM())
	}
	return bson.M{"$and": res}
}

func (f *FilterAnd) And(filter Filter) Filter {
	f.filters = append(f.filters, filter)
	return f
}

func (f *FilterAnd) Or(filter Filter) Filter {
	return Or(f, filter)
}

type FilterOr struct {
	filters []Filter
}

func (f *FilterOr) GetTp() FilterTp {
	return FilterTpOr
}

func (f *FilterOr) GetFilters() []Filter {
	return f.filters
}

func (f *FilterOr) ToBsonM() bson.M {
	res := bson.A{}
	for _, f := range f.filters {
		res = append(res, f.ToBsonM())
	}
	return bson.M{"$or": res}
}

func (f *FilterOr) And(filter Filter) Filter {
	return And(f, filter)
}

func (f *FilterOr) Or(filter Filter) Filter {
	f.filters = append(f.filters, filter)
	return f
}

type FilterIn struct {
	key   string
	value any
}

func (f *FilterIn) GetTp() FilterTp {
	return FilterTpIn
}

func (f *FilterIn) GetFilters() []Filter {
	return nil
}

func (f *FilterIn) ToBsonM() bson.M {
	return bson.M{f.key: bson.M{"$in": f.value}}
}

func (f *FilterIn) And(filter Filter) Filter {
	return And(f, filter)
}

func (f *FilterIn) Or(filter Filter) Filter {
	return Or(f, filter)
}

type FilterEq struct {
	key   string
	value any
}

func (f *FilterEq) GetTp() FilterTp {
	return FilterTpEq
}

func (f *FilterEq) GetFilters() []Filter {
	return nil
}

func (f *FilterEq) ToBsonM() bson.M {
	return bson.M{f.key: f.value}
}

func (f *FilterEq) And(filter Filter) Filter {
	return And(f, filter)
}

func (f *FilterEq) Or(filter Filter) Filter {
	return Or(f, filter)
}
