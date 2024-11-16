package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/constraints"
	"net/http"
	"strconv"
)

type HandlerFunc[R any, S any] func(ctx context.Context, req R) (resp *S, err error)
type HandlerFuncWithId[D constraints.Ordered, R any, S any] func(ctx context.Context, id D, req R) (resp *S, err error)
type HandlerFuncWithPage[R any, S any] func(ctx context.Context, page, pageSize int, req R) (resp *S, err error)

func canBind(c *gin.Context) bool {
	switch c.Request.Method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		return true
	default:
		return false
	}
}

func bind[R any](c *gin.Context, req *R) error {
	if !canBind(c) {
		return nil
	}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return err
	}
	return nil
}

func afterHandler[R any](c *gin.Context, resp *R, err error) {
	if err == nil {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"mag": err.Error()})
}

// HandlerWithJson 处理业务逻辑，参数从 JSON 格式的请求体中取
func HandlerWithJson[R any, S any](handler HandlerFunc[R, S]) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := WithGinCtx(context.Background(), c)
		req := new(R)
		if err := bind(c, req); err != nil {
			return
		}
		resp, err := handler(ctx, *req)
		afterHandler(c, resp, err)
	}
}

// HandlerWithJsonWithPage 从 URL 参数中解析 page 和 page_size 直接传递给业务 handler
//
//	 // http://localhost/api/item/list?page=2&page_size=25
//		func handler(ctx context.Context, page, pageSize int, req None, resp Resp) (resp Resp, err error) {
//		    fmt.Println(page, pageSize) // 2 25
//		    return Resp{}, nil
//		}
//		g.GET("api/item/list", HandlerWithJsonWithPage(handler))
func HandlerWithJsonWithPage[R any, S any](handler HandlerFuncWithPage[R, S]) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := WithGinCtx(context.Background(), c)
		req := new(R)
		if err := bind(c, req); err != nil {
			return
		}
		page := QueryDefaultInt(c, "page", 0)
		pageSize := QueryDefaultInt(c, "page_size", 20)
		resp, err := handler(ctx, page, pageSize, *req)
		afterHandler(c, resp, err)
	}
}

func bindAndHandle[D constraints.Ordered, R any, S any](
	ctx context.Context, c *gin.Context, handler HandlerFuncWithId[D, R, S], id D,
) {
	req := new(R)
	if err := bind(c, req); err != nil {
		return
	}
	resp, err := handler(ctx, id, *req)
	afterHandler(c, resp, err)
}

// HandlerWithUrlStr url 中携带了一个 string 类型的参数
//
//	func handler(ctx context.Context, name string, req None, resp Resp) (resp Resp, err error) {
//	    fmt.Println(name)
//	    return Resp{}, nil
//	}
//	g.POST("api/item/:name", HandlerWithUrlStr("name", handler))
func HandlerWithUrlStr[R any, S any](key string, handler HandlerFuncWithId[string, R, S]) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := WithGinCtx(context.Background(), c)
		id := c.Param(key)
		bindAndHandle[string, R, S](ctx, c, handler, id)
	}
}

// HandlerWithUrlInt url 中携带了一个 int 类型的参数
//
//	func handler(ctx context.Context, id int64, req None, resp Resp) (resp Resp, err error) {
//	    fmt.Println(id)
//	    return Resp{}, nil
//	}
//	g.POST("api/item/:id", HandlerWithUrlInt("id", handler))
func HandlerWithUrlInt[D constraints.Integer, R any, S any](
	key string, handler HandlerFuncWithId[D, R, S],
) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := WithGinCtx(context.Background(), c)
		idS := c.Param(key)
		id, err := strconv.ParseInt(idS, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "not number"})
			return
		}
		bindAndHandle(ctx, c, handler, D(id))
	}
}

type DoChildRouteFunc func(g *gin.RouterGroup) // 子路由

// RouteAdapter 用于适配 Engine 和 RouteGroup
type RouteAdapter interface {
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
}

// URLPatterns URL 调度器
func URLPatterns(route RouteAdapter, path string, childRoute DoChildRouteFunc, middles ...gin.HandlerFunc) {
	group := route.Group(path)
	for _, mid := range middles {
		group.Use(mid)
	}
	childRoute(group)
}
