package main

import (
	"WebCalculator/dal/mysql"
	"WebCalculator/router"
	"github.com/gin-gonic/gin"
)

func main() {
	mysql.Init()
	r := gin.Default()
	router.SetUpRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
