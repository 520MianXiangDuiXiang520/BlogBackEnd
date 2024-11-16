package node

import (
	"JuneBlog/patch/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type LifeCycle interface {
	PreStart() error
	Start() error
	PreStop() error
	Stop() error
}

type Node struct {
	lf       LifeCycle
	quitChan chan os.Signal
}

func NewNode(lf LifeCycle) *Node {
	return &Node{
		lf:       lf,
		quitChan: make(chan os.Signal),
	}
}

func (n *Node) Run() {
	if n.lf == nil {
		panic("Node: lifeCycle is nil")
	}

	logger.Info("Node: starting...")

	err := n.lf.PreStart()
	if err != nil {
		logger.Error("Node: PreStart fail", slog.String("err", err.Error()))
		panic(err)
	}
	logger.Info("Node: PreStart ok.")

	err = n.lf.Start()
	if err != nil {
		logger.Error("Node: Start fail", slog.String("err", err.Error()))
		panic(err)
	}
	logger.Info("Node: Start ok.")

	// wait process quit
	signal.Notify(n.quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-n.quitChan

	logger.Info("Node: start stop...")

	err = n.lf.PreStop()
	if err != nil {
		logger.Error("Node: PreStop fail", slog.String("err", err.Error()))
		return
	}
	logger.Info("Node: PreStop ok.")

	err = n.lf.Stop()
	if err != nil {
		logger.Error("Node: Stop fail", slog.String("err", err.Error()))
		return
	}
	logger.Info("Node: Stop ok. bye~")
}

/*
核心功能
  列表
    标签
  详情
  新建

其他
  sitemap
  从 github 同步
  数据存 github
  评论
  监控
  工具集

ci/cd
*/
