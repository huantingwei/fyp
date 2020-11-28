package object

/*
Deployment
	deploymentSpec
		podTemplateSpec
			podSpec
				container
					resourceRequirements
					containerPort
	deploymentStatus
		deploymentCondition
*/

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#deployment-v1-apps
type Deployment struct{
	kind               string
	objectMeta         objectMeta
	deploymentSpec     deploymentSpec
	deploymentStatus   deploymentStatus 
}

//desired behavior of the Deployment
type deploymentSpec struct{
	paused   bool              //true if the deployment is paused
	replicas int               //number of desired pods
	template podTemplateSpec
	labels   map[string]string //matchLabels, TBD
}

type deploymentStatus struct{
	availableReplicas int  //reserved and available
	readyReplicas     int  //ready for use
	replicas          int  //as specified by spec
	updatedReplicas   int  //actually operating
	conditions        deploymentCondition
}

//API doc:https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#podtemplatespec-v1-core
//desired behavior of the pod
type podTemplateSpec struct{
	objectMeta objectMeta
	podSpec    podSpec
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#podspec-v1-core
type podSpec struct{
	//array of containers inside the pod
	containers    []container
	restartPolicy string //always, on failure, or never
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#deploymentcondition-v1-apps
type deploymentCondition struct{
	message string
	status  string //True, false, or unknown
}

