package object

// Nodepool is a collection of Node
type Nodepool struct {
	//Node pool general info
	Name               string `json:"name"`
	Version            string `json:"version"`
	Location           string `json:"location"`
	Status             int    `json:"status"` //refer to NodePool_Status
	AutoscalingEnabled bool   `json:"autoscalingEnabled"`
	InitialNodeCount   int    `json:"initialNodeCount"`

	//Node pool configuration
	ImageType   string `json:"imageType"`
	MachineType string `json:"machineType"`
	DiskType    string `json:"diskType"`
	DiskSize    int    `json:"diskSize"`

	//Node pool management
	AutoUpgrade bool `json:"authUpgrade"`
	AutoRepair  bool `json:"autoRepair"`

	//Node pool security
	ServiceAccount      string `json:"serviceAccount"`
	SecureBoot          bool   `json:"secureBoot"`
	IntegrityMonitoring bool   `json:"integrityMonitoring"`
}
