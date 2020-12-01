type nodepool struct{
	//Node pool general info
	name string
	version string
	location string
	status int //refer to NodePool_Status
	autoscalingEnabled bool
	initialNodeCount int

	//Node pool configuration
	imageType string
	machineType string
	diskType string
	diskSize int

	//Node pool management
	autoUpgrade bool
	autoRepair bool

	//Node pool security
	serviceAccount string
	secureBoot bool
	integrityMonitoring bool	
}