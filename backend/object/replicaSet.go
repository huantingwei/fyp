package object

type ReplicaSet struct {
	// Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// ReplicaSet.Spec
	Replicas int `json:"Replicas"`
	// Spec.Selector.MatchLabels
	MatchLabels map[string]string `json:"Match Labels"`

	// ReplicaSet.Status
	AvailableReplicas int `json:"Available Replicas"`
	ReadyReplicas     int `json:"Ready Replicas"`
}
