package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	//创建一个数据库的连接
	var err error
	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{
		api := v1.Group("/test")
		api.POST("/", create)
		api.GET("/", fetchAll)
		api.GET("/:id", fetchSingle)
		api.PUT("/:id", update)
		api.DELETE("/:id", delete)
	}
	println("Listening and serving HTTP on 127.0.0.1:8080\n")
	router.Run()
}
