package mysql

import (
	"fmt"
	"galhub/internal/infrastructure/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() error {
	cfg := config.GlobalConfig.MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return err
	}
	DB = db
	fmt.Printf("数据库连接成功: %v\n", cfg)
	return nil
}
