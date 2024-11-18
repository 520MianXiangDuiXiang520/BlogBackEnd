package main

import (
	"JuneBlog/cmd/api"
	"JuneBlog/internal/config"
	"JuneBlog/internal/db"
	"JuneBlog/internal/env"
	"JuneBlog/internal/initialization"
	"JuneBlog/patch/logger"
	"JuneBlog/patch/node"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type JuneBlog struct {
	engine *gin.Engine
}

func (l *JuneBlog) PreStart() error {
	// init env
	env.InitGlobalEnv()

	// init logger
	err := initialization.InitLogger(env.G.LogPath)
	if err != nil {
		return err
	}

	// init config
	err = config.InitConfig(env.G.CfgPath)
	if err != nil {
		return err
	}

	// init database
	err = db.InitDatabase()
	if err != nil {
		return err
	}

	// init server
	gin.DefaultWriter = logger.GetOutput(slog.LevelInfo)
	gin.DefaultErrorWriter = logger.GetOutput(slog.LevelError)
	l.engine = gin.Default()
	api.RegisterRouter(l.engine)
	return nil
}

func (l *JuneBlog) Start() error {
	addr := config.G.GetStrWithDefault(config.CommonKeyAddr, ":8080")
	cert := config.G.GetStrWithDefault(config.CommonKeyTLSCertFile, "")
	key := config.G.GetStrWithDefault(config.CommonKeuTLSKeyFile, "")
	if cert == "" || key == "" {
		return l.engine.Run(addr)
	}
	return l.engine.RunTLS(addr, cert, key)
}

func (l *JuneBlog) PreStop() error {
	return nil
}

func (l *JuneBlog) Stop() error {
	return nil
}

func main() {
	instance := node.NewNode(&JuneBlog{})
	instance.Run()
}
