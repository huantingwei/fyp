package interface

type Pod struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Spec        *PodSpec
	Status      *PodStatus
}

type PodTemplateSpec struct {
	Spec   *PodSpec
	Status *PodStatus
}

type PodSpec struct {
	Containers         []*Container
	SecurityContext    *PodSecurityContext
	ServiceAccountName string
	Volumes            []*Volume
	NodeSelector       map[string]string
	// some other...
}
type PodStatus struct {
	Conditions      []*PodCondition
	ContainerStatus []*ContainerStatus
	HostIP          string
	PodIP           string
	PodIPs          []string
	Phase           string
	Message         string
}

type PodSecurityContext struct {
	// some complicated stuff
}

type PodCondition struct {
	Message string
	Reason  string
	Staus   string
	Type    string
}
