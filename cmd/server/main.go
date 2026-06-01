package main

import (
	"fmt"
	"galhub/internal/infrastructure/config"
	"log"
)

func main() {
	if err := config.LoadConfig("configs/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("配置初始化成功: %v\n", config.GlobalConfig.Server)
}
