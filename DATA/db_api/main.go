package main

import (
	test "db_api/src/test"
	user "db_api/src/user"

	docs "db_api/docs"
	// docs "db_api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// @title DataBase API
// @version 1.0
func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1/")

	test.Router(v1, "test")
	user.Router(v1, "user")

	router.GET("/docs/*any", gs.WrapHandler(swaggerfiles.Handler))
	println("\nListening and serving HTTP on http://127.0.0.1:8080\n")
	router.Run()
}
