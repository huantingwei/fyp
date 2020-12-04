package object

type Nodepool struct{
	//Node pool general info
	Name string
	Version string
	Location string
	Status int //refer to NodePool_Status
	AutoscalingEnabled bool
	InitialNodeCount int

	//Node pool configuration
	ImageType string
	MachineType string
	DiskType string
	DiskSize int

	//Node pool management
	AutoUpgrade bool
	AutoRepair bool

	//Node pool security
	ServiceAccount string
	SecureBoot bool
	IntegrityMonitoring bool	
}