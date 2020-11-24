package interface

type VersionKind struct {
	ApiVersion string
	Kind       string
}
type ObjectMeta struct {
	Name        string
	ClusterName string
	Namespace   string
	Labels      map[string]string
}

type LabelSelector struct {
	MatchLabels      map[string]string
	MatchExpressions []*LabelSelectorRequirement
}

type LabelSelectorRequirement struct {
	Key      string
	Operator string
	Values   string
}
