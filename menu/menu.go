package menu

import (
	"github.com/gin-gonic/gin"
	"myProject/dataConfig"
	db "myProject/gorm"
	"net/http"
)

func CreateMenu(c *gin.Context) {
	// 解析客户端传来的 JSON 数据
	var menu dataConfig.Menu
	err := c.BindJSON(&menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建菜单记录
	if err = db.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	return
}
