package object

// Pod is the minimum execution unit
type Pod struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	//object metadata - extra
	Labels map[string]string `json:"Labels"`
	//details of the pod owner, e.g. deployment
	OwnerReferences []OwnerReference `json:"Owner References"`
	ControlledBy    string           `json:"Controlled By"`

	//Pod spec
	Containers    []Container `json:"Containers"`
	DNSPolicy     string      `json:"DNS Policy"`
	RestartPolicy string      `json:"Restart Policy"`
	NodeName      string      `json:"Node Name"`

	//Pod status
	HostIP            string            `json:"Host IP"`
	PodIP             string            `json:"Pod IP"`
	Phase             string            `json:"Phase"` //pending, running, failed, etc
	ContainerStatuses []ContainerStatus `json:"Container Statuses"`
}

// OwnerReference specify what owns the Pod
type OwnerReference struct {
	Name string `json:"Name"`
	Uid  string `json:"UID"`
	Kind string `json:"Kind"`
}

// ContainerStatus describes the status of containers
type ContainerStatus struct {
	Name         string `json:"Name"`
	Image        string `json:"Image"`
	Ready        bool   `json:"Ready"`
	RestartCount int    `json:"Restart Count"`
}
