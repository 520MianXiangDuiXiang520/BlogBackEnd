package message

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	RespHeader
	Token string `json:"token"`
}
