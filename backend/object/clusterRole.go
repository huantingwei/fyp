package object

// ClusterRole
type ClusterRole struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// ClusterRole spec
	// AggregationRule.ClusterRoleSelectors.MatchLabels
	ClusterRoleSelectors map[string]string `json:"Cluster Role Selectors"`
	Rules                []PolicyRule      `json:"Rules"`
}
