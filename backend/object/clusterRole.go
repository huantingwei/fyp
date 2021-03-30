package object

// ClusterRole
type ClusterRole struct {
	ObjectMeta 				ObjectMeta 			`json:"objectMeta"`

	// ClusterRole spec
	// AggregationRule.ClusterRoleSelectors.MatchLabels
	ClusterRoleSelectors 	map[string]string 	`json:"clusterRoleSelectors"`
	Rules      				[]PolicyRule		`json:"rules"`
}
