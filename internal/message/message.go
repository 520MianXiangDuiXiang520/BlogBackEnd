package message

import "net/http"

type None struct{}

type PagerReq struct {
	Page     int64 `json:"page" check:"more: 0"`
	PageSize int64 `json:"page_size" check:"not null; size: [5,20]"`
}

type RespHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	RespOK = RespHeader{Code: http.StatusOK, Msg: "ok"}
)
