package history

import (
	"WebCalculator/dal/mysql"
	"WebCalculator/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddHistory(c *gin.Context) {
	var history entity.History
	err := c.ShouldBindJSON(&history)
	if err != nil {
		panic(err)
	}
	mysql.DB.Create(&history)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
	})
}
