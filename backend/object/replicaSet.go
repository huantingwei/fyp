package object

type ReplicaSet struct {
	// Object metadata
	ObjectMeta ObjectMeta 					`json:"objectMeta"`

	// ReplicaSet.Spec
	Replicas			int					`json:"replicas"`
	// Spec.Selector.MatchLabels
	MatchLabels 		map[string]string	`json:"matchLabels"`

	// ReplicaSet.Status
	AvailableReplicas	int					`json:"availableReplicas"`
	ReadyReplicas       int 				`json:"readyReplicas"`
}

