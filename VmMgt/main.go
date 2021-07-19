package main

import (
	"github.com/gin-gonic/gin"
	"ssdd.com/vms/controllers"
)

func main() {
	r := gin.Default()
	r.Use(AssignServiceId())
	r.Use(Audit())
	r.Use(Auth())
	r.POST("/vms", controllers.Register)
	r.GET("/vms/:name", controllers.Get)
	r.GET("/vms", controllers.List)
	r.PATCH("/vms/:name", controllers.Update)
	r.DELETE("/vms/:name", controllers.Delete)
	r.Run(":8080")
}
