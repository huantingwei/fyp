package object

type StatefulSet struct {
	// Object metadata
	ObjectMeta ObjectMeta 					`json:"objectMeta"`

	// StatefulSet.Spec
	Replicas			int					`json:"replicas"`
	// Spec.Selector.MatchLabels
	MatchLabels 		map[string]string	`json:"matchLabels"`
	// Service that governs this StatefulSet
	ServiceName			string				`json:"serviceName"`
	PodManagementPolicy	string				`json:"podManagementPolicy"`

	// StatefulSet.Status
	CurrentReplicas		int					`json:"currentReplicas"`
	UpdatedReplicas     int 				`json:"updatedReplicas"`
	ReadyReplicas       int 				`json:"readyReplicas"`
}

