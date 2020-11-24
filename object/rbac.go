package object

type Subject struct {
	ApiGroup  *string
	Kind      *string
	Name      *string
	NameSpace *string
}

type PolicyRule struct {
	ApiGroups       []*string
	ResourceNames   []*string
	Resources       []*string
	Verbs           []*string
	NonResourceURLs []*string
}

type RoleRef struct {
	ApiGroups []*string
	Kind      *string
	Name      *string
}

type Role struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Rules       []*PolicyRule
}

type RoleBinding struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	RoleRef     *RoleRef
	Subject     []*Subject
}
