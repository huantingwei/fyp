package object

type StatefulSet struct {
	// Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// StatefulSet.Spec
	Replicas int `json:"Replicas"`
	// Spec.Selector.MatchLabels
	MatchLabels map[string]string `json:"Match Labels"`
	// Service that governs this StatefulSet
	ServiceName         string `json:"Service Name"`
	PodManagementPolicy string `json:"Pod Management Policy"`

	// StatefulSet.Status
	CurrentReplicas int `json:"Current Replicas"`
	UpdatedReplicas int `json:"Updated Replicas"`
	ReadyReplicas   int `json:"Ready Replicas"`
}
