package object

// Cluster of Kubernetes
type Cluster struct {
	//Cluster general info
	Name           string `json:"name"`
	CreationTime   string `json:"creationTime"`
	MasterVersion  string `json:"masterVersion"`
	IPendpoint     string `json:"ipEndpoint"`
	Location       string `json:"location"`
	ReleaseChannel int    `json:"releaseChannel"`
	Status         string `json:"status"`

	//Cluster networking config
	Network                  string `json:"network"`
	NetworkConfig            string `json:"networkConfig"`
	Subnet                   string `json:"subnet"`
	IntranodeVisibility      bool   `json:"intranodeVisibility"`
	NetworkPolicyEnabled     bool   `json:"networkPolicyEnabled"`
	MasterAuthNetworkEnabled bool   `json:"materAuthNetworkEnabled"`

	//Cluster security config
	ShieldedNodeEnabled        bool `json:"shieldNodeEnabled"`
	BinaryAuthorisationEnabled bool `json:"binaryAuthorisationEnabled"`
	ClientCertificateEnabled   bool `json:"clientCertificateEnabled"`
}
