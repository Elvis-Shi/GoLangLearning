package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	MachineName  string
	ImageName    string
	CpuCores     int
	MemorySizeMB int
	Region       int // region id.
}

type UpdateStatusInput struct {
	MachineName string
	Status      string
}

type GetInput struct {
	MachineName string
}

type DeleteInput struct {
	MachineName string
}

type ListAllInput struct {
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storage := CacheStorage{}
	vm, err := storage.Register(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vm)
}

func Get(c *gin.Context) {
	// check if exists

	// retrieve from storage.
}

func UpdateStatus(c *gin.Context) {
	// check if exists

	// check if current status available for update to the specified new status.
}

func ListAll(c *gin.Context) {

}

func Delete(c *gin.Context) {
	// check if exists

	// check status ok for deletion.

	// delete VM.

	// update storage
}
