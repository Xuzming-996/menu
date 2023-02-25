package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myProject/dataConfig"
)

var DB *gorm.DB

func NewDBConn() {
	dsn := "root:root@tcp(localhost:3306)/menu"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = db
	CheckTable()
}

func CheckTable() {
	//连接菜单数据库
	migrator := DB.Migrator()
	// 检查 Menu 表是否存在
	if !migrator.HasTable(&dataConfig.Menu{}) {
		// 如果不存在，则创建 Menu 表
		err := migrator.CreateTable(&dataConfig.Menu{})
		if err != nil {
			panic("failed to create table")
		}
	}
}
