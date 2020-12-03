package object

type Cluster struct{
	//Cluster general info
	Name           string
	CreationTime   string 
	MasterVersion  string
	IPendpoint     string 
	Location       string
	ReleaseChannel int
	Status         string

	//Cluster networking config
	Network                  string
	NetworkConfig            string
	Subnet                   string
	IntranodeVisibility      bool
	NetworkPolicyEnabled     bool
	MasterAuthNetworkEnabled bool

	//Cluster security config
	ShieldedNodeEnabled        bool
	BinaryAuthorisationEnabled bool
	ClientCertificateEnabled   bool
}

