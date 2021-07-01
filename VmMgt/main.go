package main

import (
	"github.com/gin-gonic/gin"
	"ssdd.com/vms/controllers"
)

func main() {
	r := gin.Default()

	r.POST("/vms", controllers.Register)
	r.GET("/vms/:name", controllers.Get)
	r.GET("/vms", controllers.ListAll)
	r.PATCH("/vms/:operation", controllers.UpdateStatus)
	r.DELETE("/vms/:name", controllers.Delete)
	r.Run()
}
