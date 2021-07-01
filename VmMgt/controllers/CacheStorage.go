package controllers

import (
	"errors"
	"fmt"

	"ssdd.com/vms/models"
)

type VMStorage interface {
	Register(input *RegisterInput) (models.VM, error) // TODO: is there any clever way instead of define all these different RegisterInput/Output?
	Get(machineName string) (models.VM, error)
	Delete(machineName string) error // TODO: what should Delete return?
	ListAll() error                  // TODO: what is arry in Go.
	UpdateStatus(machineName string, status string) (models.VM, error)
}

var vmsCache = make(map[string]*models.VM)
var imagesCache = map[string]*models.VMImage{
	"Ubuntu 14.04": &models.VMImage{ImageName: "Ubuntu 14.04", OS: "Ubuntu", Version: "14.04", SizeMB: 4000},
}

type CacheStorage struct {
}

func (storage CacheStorage) Register(input *RegisterInput) (*models.VM, error) {
	// TODO: check authentication/authorization.
	// TODO: log activity?

	if input == nil {
		return nil, errors.New("register input is nil")
	}

	if _, ok := vmsCache[input.MachineName]; ok {
		// machine name used.

		return nil, fmt.Errorf("machine name %w already used. Please choose a different one", input.MachineName)
	}

	image, ok := imagesCache[input.ImageName]
	if !ok {
		// image invalid.

		return nil, fmt.Errorf("image %w not exists!", input.ImageName)
	}

	// more checks,
	// check region limitation
	// privision resource(CPU, memory in subscription) available to satisfy required.

	// Start VM privision in other thread.

	vm := models.VM{
		MachineName:  input.MachineName,
		Image:        *image,
		IP:           "", // TODO: update when privision job done.
		CpuCores:     input.CpuCores,
		MemorySizeMB: input.MemorySizeMB,
		Region:       input.Region,
		Status:       "Privisioning",
	}
	return &vm, nil
}
