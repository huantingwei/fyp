package object

// Deployment is a set of Pods with template
type Deployment struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	//deployment spec
	DesiredPods int               `json:"Desired Pods"`
	MatchLabels map[string]string `json:"Match Labels"`

	//Pod spec for this deployment
	Containers    []Container `json:"Containers"`
	DNSPolicy     string      `json:"DNS Policy"`
	RestartPolicy string      `json:"Restart Policy"`
	NodeName      string      `json:"Node Name"`

	//deployment status
	UpdatedReplicas     int `json:"Updated Replicas"`
	ReadyReplicas       int `json:"Ready Replicas"`
	AvailableReplicas   int `json:"Available Replicas"`
	UnavailableReplicas int `json:"Unavailable Replicas"`
}
