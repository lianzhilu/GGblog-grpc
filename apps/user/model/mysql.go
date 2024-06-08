package model

import (
	"context"
	"fmt"
	"ggblog-grpc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDataBase() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLConf.User,
		config.MySQLConf.Password,
		config.MySQLConf.Port,
		config.MySQLConf.Name)
	fmt.Println("dsn", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to database failed", err.Error())
	}

	// db.AutoMigrate(&User{}, &Article{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// sqlDB.SetConnMaxLifetime(time.Hour)
}

func AddContext(ctx context.Context) *gorm.DB {
	tmpDB := db
	return tmpDB.WithContext(ctx)
}
