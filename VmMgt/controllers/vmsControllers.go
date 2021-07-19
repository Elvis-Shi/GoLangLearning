package controllers

import (
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

func Update(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Machine name is required."})
		return
	}

	var input models.UpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vm, err := storage.Update(name, &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}
