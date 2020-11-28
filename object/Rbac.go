package object

//'kind' means the kind of REST resource a struct represents

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#roleref-v1-rbac-authorization-k8s-io
//info of a cluster role
type roleRef struct{
	kind string `default: "RoleRef"`
	name string
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#subject-v1-rbac-authorization-k8s-io
//the object where a role is binding to
type subject struct{
	kind      string `default: "Subject"`
	name      string
	namespace string
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#clusterrolebinding-v1-rbac-authorization-k8s-io
//cluster role is a role that applies to the entire cluster, rather than a namespace
//cluster role can grant cluster-level permission to users
type ClusterRoleBinding struct{
	kind string `default: "ClusterRoleBinding"`
	objectMeta objectMeta
	roleRef roleRef
	subjects []subject


}