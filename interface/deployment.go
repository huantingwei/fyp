package interface

type Deployment struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Spec        *DeploymentSpec
	Status      *DeploymentStatus
}

type DeploymentSpec struct {
	Replicas int
	Selector *LabelSelector
	Strategy *DeploymentStrategy
	Template *PodTemplateSpec
}

type DeploymentStatus struct {
	Conditions          []*DeploymentCondition
	Replicas            int
	AvailableReplicas   int
	ReadyReplicas       int
	UnavailableReplicas int
	UpdatedReplicas     int
}

type DeploymentStrategy struct {
	Type          string // RollingUpdate(default) or Recreate
	RollingUpdate *RollingUpdateDeployment
}

type DeploymentCondition struct {
	Type    string
	Status  string
	Message string
	Reason  string
}

type RollingUpdateDeployment struct {
	MaxSurge       int
	MaxUnavailable int
}
