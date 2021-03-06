package main

import (
	"github.com/Guiyunweb/shiki/conf"
	"report-api/cmd"
)

func main() {
	// 读取配置
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// 启动服务
	cmd.Run()
}
