type cluster struct{
	//Cluster general info
	name string
	createTime string 
	masterVersion string
	IPendpoint string
	location string
	releaseChannel int
	status string

	//Cluster networking config
	network string
	networkConfig string
	subnet string
	intranodeVisibility bool
	networkPolicyEnabled bool
	masterAuthNetworkEnabled bool

	//Cluster security config
	shieldedNodeEnabled bool
	binaryAuthorisationEnabled bool
	clientCertificateEnabled bool
}