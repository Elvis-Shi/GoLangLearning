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
	List() error                     // TODO: what is arry in Go.
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

		return nil, fmt.Errorf("machine name %v already used. Please choose a different one", input.MachineName)
	}

	image, ok := imagesCache[input.ImageName]
	if !ok {
		// image invalid.

		return nil, fmt.Errorf("image %v not exists", input.ImageName)
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

	vmsCache[input.MachineName] = &vm

	return &vm, nil
}

func (storage CacheStorage) Get(machineName string) (*models.VM, error) {
	// check if exists
	vm, ok := vmsCache[machineName]
	if !ok {
		return nil, fmt.Errorf("machine name %v not exists", machineName)
	}

	return vm, nil
}

func (storge CacheStorage) List() ([]*models.VM, error) {
	vms := make([]*models.VM, 0, len(vmsCache))
	for _, value := range vmsCache {
		vms = append(vms, value)
	}

	return vms, nil
}

func (storage CacheStorage) Delete(machineName string) (*models.VM, error) {
	// check if exists
	vm, ok := vmsCache[machineName]
	if !ok {
		return nil, fmt.Errorf("machine name %v not exists", machineName)
	}

	// TODO: check any vm deletion restriction.
	// TODO: delete vm async.
	// TODO: whether to delete VM disk?

	vm.Status = "Deleting"

	return vm, nil
}

func (storage CacheStorage) Operate(machineName string, operation string) (*models.VM, error) {
	// check if exists
	vm, ok := vmsCache[machineName]
	if !ok {
		return nil, fmt.Errorf("machine name %v not exists", machineName)
	}

	if operation != "start" && operation != "shutdown" && operation != "reboot" {
		return nil, fmt.Errorf("operation %v is not allowed, allowed operations are start, shutdown, reboot", operation)
	}

	switch vm.Status {
	case "Provisioning", "Shutingdown", "Deleting", "Deleted":
		return nil, fmt.Errorf("cannot update status for VM %v as its status is currently %v", machineName, vm.Status)
	case "Shutdown":
		if operation == "shutdown" {
			return vm, nil
		} else if operation == "start" {
			// TODO: start VM async.
			vm.Status = "Running"
			return vm, nil
		} else {
			return nil, fmt.Errorf("cannot reboot machine while its status is %v", vm.Status)
		}
	case "Running":
		if operation == "start" {
			return vm, nil // do nothing
		} else if operation == "shutdown" {
			// TODO: stop VM async.
			vm.Status = "Shuttingdown"
			return vm, nil
		} else {
			// TODO: reboot VM async.
			vm.Status = "Shuttingdown"
			return vm, nil
		}
	default:
		return nil, fmt.Errorf("machine is in invalid status %v", vm.Status)
	}
}
