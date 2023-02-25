package main

import (
	"github.com/gin-gonic/gin"
	db "myProject/gorm"
	"myProject/menu"
)

func initRouter() *gin.Engine {
	r := gin.Default()

	//传入菜单
	r.POST("/create", menu.CreateMenu)
	return r
}

func main() {
	db.NewDBConn()
	r := initRouter()
	r.Run(":8080")
}
