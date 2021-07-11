package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"ssdd.com/vms/models"
)

var storage = models.CacheStorage{}

func Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vm, err := storage.Register(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}

func Get(c *gin.Context) {
	name := c.Param("name")
	// validation not required, since if name not provided, List API will be invoked instead.

	vm, err := storage.Get(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
	
}

func List(c *gin.Context) {
	vms, err := storage.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vms)
}

func Delete(c *gin.Context) {
	name := c.Param("name")
	vm, err := storage.Delete(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}

func Operate(c *gin.Context) {
	name := c.Param("name")
	operation := c.Param("operation")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Machine name is required."})
		return
	}

	if operation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation is required"})
		return
	}

	if operation != "start" && operation != "shutdown" && operation != "reboot" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Operation %v is not allowed.", operation)})
	}

	vm, err := storage.Operate(name, operation)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}
