package object

// Cluster of Kubernetes
type Cluster struct {
	//Cluster general info
	Name           string `json:"Name"`
	CreationTime   string `json:"Creation Time"`
	MasterVersion  string `json:"Master Version"`
	IPendpoint     string `json:"IP Endpoint"`
	Location       string `json:"Location"`
	ReleaseChannel int    `json:"Release Channel"`
	Status         string `json:"Status"`

	//Cluster networking config
	Network                  string `json:"Network"`
	NetworkConfig            string `json:"Network Config"`
	Subnet                   string `json:"Subnet"`
	IntranodeVisibility      bool   `json:"IntraNode Visibility"`
	NetworkPolicyEnabled     bool   `json:"Network Policy Enabled"`
	MasterAuthNetworkEnabled bool   `json:"MasterAuth Network Enabled"`

	//Cluster security config
	ShieldedNodeEnabled        bool `json:"Shield Node Enabled"`
	BinaryAuthorisationEnabled bool `json:"Binary Authorisation Enabled"`
	ClientCertificateEnabled   bool `json:"Client Certificate Enabled"`
}
