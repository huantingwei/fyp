package interface

type ResourceRequirements struct {
	Limits   map[string]string
	Requests map[string]string
}
type ContainerPort struct {
	Name          string
	Protocol      string
	ContainerPort int
	HostIP        string
	HostPort      int
}

type SecurityContext struct {
}

type Container struct {
	Name            string
	Args            []string
	Command         []string
	Env             interface{}
	EnvFrom         interface{}
	Image           string
	ImagePullPolicy string
	Ports           []*ContainerPort
	Resources       *ResourceRequirements
	SecurityContext *SecurityContext
}

type ContainerStatus struct {
	Name        string
	ContainerID string
	Image       string
	ImageID     string
	Ready       bool
	State       string // either RUNNING, TERMINATED, or WAITING
}

const (
	RUNNING    = "running"
	TERMINATED = "terminated"
	WAITING    = "waiting"
)
