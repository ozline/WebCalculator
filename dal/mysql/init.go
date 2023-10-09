package mysql

import (
	"WebCalculator/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "root:aaa24331277@tcp(localhost:3306)/calculator?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&entity.History{}, &entity.Deposit{}, &entity.Loan{})
	if err != nil {
		return
	}
}
