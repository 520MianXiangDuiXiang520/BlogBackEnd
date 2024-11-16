package opt

import "go.mongodb.org/mongo-driver/bson"

type Ctx struct {
	Page     int64
	PageSize int64
	Filter   Filter
	OrderBy  string
}

func (c *Ctx) ApplyFilter() bson.M {
	if c.Filter == nil {
		return bson.M{}
	}
	return c.Filter.ToBsonM()
}

func NewAndApplyCtx(opts ...Opt) *Ctx {
	ctx := &Ctx{
		OrderBy: "id",
	}
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

type Opt func(ctx *Ctx)

func WithOrderBy(field string) Opt {
	return func(ctx *Ctx) {
		ctx.OrderBy = field
	}
}

func WithPage(n int64) Opt {
	return func(ctx *Ctx) {
		ctx.Page = n
	}
}

func WithPageSize(n int64) Opt {
	return func(ctx *Ctx) {
		ctx.PageSize = n
	}
}

func WithFilter(f Filter) Opt {
	return func(ctx *Ctx) {
		if ctx.Filter == nil {
			ctx.Filter = f
			return
		}
		ctx.Filter.And(f)
	}
}
