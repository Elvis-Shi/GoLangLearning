package models

type VM struct {
	MachineName  string // Primary key for now.
	Image        VMImage
	IP           string
	CpuCores     int
	MemorySizeMB int
	Region       int // region id.
	// allowed values: Runing, Shutingdown/Shutdown, Provisioning, Deleting/Deleted(in case some related data still required).
	Status string // TODO: try to use enum like data instead of string.
}

type VMImage struct {
	ImageName string // Primary key for Image.
	OS        string
	Version   string
	SizeMB    int
}

type RegisterInput struct {
	MachineName  string
	ImageName    string
	CpuCores     int
	MemorySizeMB int
	Region       int // region id.
}
