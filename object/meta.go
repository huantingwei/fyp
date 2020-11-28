package object

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#ownerreference-v1-meta
//e.g. a Deployment generated a pod, so the Deployment is its owner
type ownerReference struct{
    name        string
    kind        string
    uid         string
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#objectmeta-v1-meta
type objectMeta struct{
    labels               map[string]string
    name                 string
    namespace            string
    ownerReference       ownerReference
    uid                  string
}