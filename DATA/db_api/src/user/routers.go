package user

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup, name string) {
	router := r.Group(name)
	{
		router.GET("/", getting)
		router.POST("/", posting)
		router.PUT("/", putting)
		router.DELETE("/", deleting)
	}
}
