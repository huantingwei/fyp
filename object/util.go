package object

// Container is used by pod and deployment
type Container struct {
	Name            string         `json:"name"`
	Image           string         `json:"image"`
	ImagePullPolicy string         `json:"imagePullPolicy"`
	ContainerPorts  map[int]string `json:"containerPorts"`
}

// ObjectMeta is used by deployment, node, pod, service
type ObjectMeta struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Uid          string `json:"uid"`
	CreationTime string `json:"creationTime"`
}
