package loans

import (
	"WebCalculator/dal/mysql"
	"WebCalculator/entity"
	"WebCalculator/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateLoans(c *gin.Context) {
	var loan entity.Loan
	err := c.ShouldBindJSON(&loan)
	if err != nil {
		panic(err)
	}
	if !service.ExistDuration(loan.Duration, 0) {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "没有对应的时长",
		})
	} else {
		mysql.DB.Where("duration = ?", loan.Duration).Updates(&loan)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	}
}
