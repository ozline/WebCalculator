package deposit

import (
	"WebCalculator/entity"
	"WebCalculator/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryInterest(c *gin.Context) {
	cal := entity.MoneyCal{}
	err := c.ShouldBindJSON(&cal)
	if err != nil {
		panic(err)
	}
	if !service.ExistDuration(cal.Duration, 1) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "不存在对应时长",
		})
		return
	}
	interest := service.CalInterest(cal, 1)
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "计算成功",
		"result": interest,
	})

}
