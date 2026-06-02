package main

import (
	"fmt"
	"galhub/internal/infrastructure/config"
	"galhub/internal/infrastructure/persistence/mysql"
	"galhub/internal/interfaces/router"
	"log"
)

func main() {
	fmt.Printf("配置初始化\n")
	if err := config.LoadConfig("configs/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("数据库初始化\n")
	if err := mysql.Init(); err != nil {
		log.Fatalf("Failed to init mysql: %v", err)
	}
	fmt.Printf("路由初始化\n")
	r := router.NewRouter()

	addr := fmt.Sprintf(
		":%d",
		config.GlobalConfig.Server.Port,
	)
	fmt.Printf("服务器启动: %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}

}
