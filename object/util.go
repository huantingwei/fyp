package object

//used by pod and deployment
type Container struct{
	Name string
	Image string
	ImagePullPolicy string
	ContainerPorts map[int]string
}

//used by deployment, node, pod, service
type ObjectMeta struct{
	Name string
	Namespace string
	Uid string
	CreationTime string
}