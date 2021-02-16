package object

// Deployment is a set of Pods with template
type Deployment struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"objectMeta"`

	//deployment spec
	DesiredPods int               `json:"desiredPods"`
	MatchLabels map[string]string `json:"matchLabels"`

	//Pod spec for this deployment
	Containers    []Container `json:"containers"`
	DNSPolicy     string      `json:"DNSPolicy"`
	RestartPolicy string      `json:"restartPolicy"`
	NodeName      string      `json:"nodeName"`

	//deployment status
	UpdatedReplicas     int `json:"updatedReplicas"`
	ReadyReplicas       int `json:"readyReplicas"`
	AvailableReplicas   int `json:"availableReplicas"`
	UnavailableReplicas int `json:"unavailableReplicas"`
}
