package object

// Nodepool is a collection of Node
type Nodepool struct {
	//Node pool general info
	Name               string `json:"Name"`
	Version            string `json:"Version"`
	Location           string `json:"Location"`
	Status             int    `json:"Status"` //refer to NodePool_Status
	AutoscalingEnabled bool   `json:"Autoscaling Enabled"`
	InitialNodeCount   int    `json:"Initial Node Count"`

	//Node pool configuration
	ImageType   string `json:"Image Type"`
	MachineType string `json:"Machine Type"`
	DiskType    string `json:"Disk Type"`
	DiskSize    int    `json:"Disk Size"`

	//Node pool management
	AutoUpgrade bool `json:"Auth Upgrade"`
	AutoRepair  bool `json:"Auto Repair"`

	//Node pool security
	ServiceAccount      string `json:"Service Account"`
	SecureBoot          bool   `json:"Secure Boot"`
	IntegrityMonitoring bool   `json:"Integrity Monitoring"`
}
