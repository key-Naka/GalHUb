package main

import (
	"fmt"
	"galhub/internal/bootstrap"
	"galhub/internal/infrastructure/config"
	"galhub/internal/infrastructure/mysql"
	"galhub/internal/interfaces/http/router"
	"galhub/internal/pkg/logger"
)

func main() {
	logger.Info("配置初始化")
	if err := config.LoadConfig("configs/config.yaml"); err != nil {
		logger.Error(fmt.Sprintf("Failed to load config: %v", err.Error()))
	}

	logger.Info("数据库初始化")
	if err := mysql.Init(); err != nil {
		logger.Error(fmt.Sprintf("Failed to init mysql: %v", err.Error()))
	}
	logger.Info("路由初始化")
	bs := bootstrap.New()

	r := router.NewRouter(bs)

	addr := fmt.Sprintf(
		":%d",
		config.GlobalConfig.Server.Port,
	)
	logger.Info(fmt.Sprintf("服务器启动: %s", addr))
	if err := r.Run(addr); err != nil {
		logger.Error(fmt.Sprintf("Failed to run server: %v", err.Error()))
	}

}
