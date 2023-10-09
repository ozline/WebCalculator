package service

import (
	"WebCalculator/dal/mysql"
	"WebCalculator/entity"
)

func GetLast10(histories []*entity.History) {
	mysql.DB.Order("created_at desc").Limit(10).Find(&histories)
}
