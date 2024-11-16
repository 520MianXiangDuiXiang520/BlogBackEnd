package logic

import (
	"JuneBlog/internal/common"
	"JuneBlog/internal/config"
	"JuneBlog/internal/db"
	"JuneBlog/internal/message"
	"JuneBlog/internal/utils"
	"JuneBlog/patch/ginx"
	"JuneBlog/patch/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckPermitted(ctx *gin.Context) bool {
	token, err := ctx.Cookie("SESSION_ID")
	if err != nil {
		return false
	}
	sToken, err := db.Db().GetToken(context.Background())
	if err != nil {
		return false
	}
	return token == sToken
}

func LoginReq(ctx context.Context, req message.LoginReq) (resp *message.LoginResp, err error) {
	psd := config.G.GetStrWithDefault(config.CommonKeyAdminPassword, "")
	if psd == "" {
		logger.Error("Cfg: password not set")
		return nil, common.SrvErrorCfg
	}
	password := utils.Sha256(req.Password)
	if password != psd {
		logger.Info("Auth: bad password", "input", req.Password, "psw", password, "p", psd)
		return nil, common.UserErrorBadPasswordOrUserName
	}
	username := config.G.GetStrWithDefault(config.CommonKeyAdminUsername, "")
	if username != "" && req.Username != username {
		logger.Info("Auth: bad username", "username", req.Username)
		return nil, common.UserErrorBadPasswordOrUserName
	}
	token := utils.UUID()
	err = db.Db().SetToken(ctx, token)
	if err != nil {
		logger.Error("Auth: SetToken", "token", token, "err", err)
		return nil, err
	}
	ginCtx := ginx.GinCtx(ctx)
	ginCtx.SetSameSite(http.SameSiteNoneMode)
	ginCtx.SetCookie("SESSION_ID", token, 3600, "/",
		ginCtx.GetHeader("Origin"), true, false)
	resp = &message.LoginResp{
		RespHeader: message.RespOK,
		Token:      token,
	}
	return resp, nil
}

func LogoutReq(ctx context.Context, _ message.None) (resp *message.RespHeader, err error) {
	err = db.Db().DelToken(ctx)
	if err != nil {
		logger.Error("Auth: DelToken", "err", err)
		return nil, err
	}
	return &message.RespOK, nil
}
