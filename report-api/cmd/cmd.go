package cmd

import (
	"github.com/Guiyunweb/shiki/log"
	"report-api/model"
	"report-api/server"
)

func Run() {
	// 数据库初始化
	log.Init()
	//初始化数据库
	model.Init()
	// 启动服务
	server.StartServer()
}

