package interface

type ConfigMap struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Data        map[interface{}]interface{}
	BinaryData  map[interface{}]interface{}
	Immutable   bool
}

type Secret struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Type        string
	Data        map[interface{}]interface{}
	StringData  string
	Immutable   bool
}
