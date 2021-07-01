package models

type VM struct {
	MachineName  string // Primary key for now.
	Image        VMImage
	IP           string
	CpuCores     int
	MemorySizeMB int
	Region       int // region id.
	// allowed values: Runing, Shutdown, Provisioning, Deleted(in case some related data still required).
	Status string // TODO: try to use enum like data instead of string.
}

type VMImage struct {
	ImageName string // Primary key for Image.
	OS        string
	Version   string
	SizeMB    int
}
