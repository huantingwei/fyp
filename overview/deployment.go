package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshDeploymentInfo(c *gin.Context){
	deploymentInfo := s.initDeploymentArray();

	_, err := s.deploymentCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	_, err2 := s.deploymentCollection.InsertMany(context.TODO(),deploymentInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
	}

	fmt.Println("refreshed deployment info")
}

func (s *Service) initDeploymentArray() []interface{}{
	deploymentList, err := s.clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var deploymentSlice []interface{};

	for _, d := range deploymentList.Items{
		deployment := object.Deployment{
			ObjectMeta: object.ObjectMeta{
				Name: d.Name,
				Namespace: string(d.Namespace),
				Uid: string(d.UID),
				CreationTime: d.CreationTimestamp.String(),
			},
			DesiredPods: int(*d.Spec.Replicas),
			DnsPolicy: string(d.Spec.Template.Spec.DNSPolicy),
			RestartPolicy: string(d.Spec.Template.Spec.RestartPolicy),
			NodeName: d.Spec.Template.Spec.NodeName,
			UpdatedReplicas: int(d.Status.UpdatedReplicas),
			ReadyReplicas: int(d.Status.ReadyReplicas),
			AvailableReplicas: int(d.Status.AvailableReplicas),
			UnavailableReplicas: int(d.Status.UnavailableReplicas),
		}

		matchLabelsMap := make(map[string]string);
		var containersSlice []object.Container;

		matchLabels := d.Spec.Selector.MatchLabels;
		for key, val := range matchLabels{
			matchLabelsMap[key] = val;
		}

		containers := d.Spec.Template.Spec.Containers;
		for _, c := range containers{
			container := object.Container{
				Name: c.Name,
				Image: c.Image,
				ImagePullPolicy: string(c.ImagePullPolicy),
			}

			containerPort := make(map[int]string);

			for _, port := range c.Ports{
				containerPort[int(port.ContainerPort)] = string(port.Protocol);
			}
			container.ContainerPorts = containerPort;
			containersSlice = append(containersSlice, container);
		}

		deployment.MatchLabels = matchLabelsMap;
		deployment.Containers = containersSlice;
		deploymentSlice = append(deploymentSlice, deployment)
	}

	return deploymentSlice;
}

func (s *Service) GetDeploymentInfo(c *gin.Context) {
	cursor, err := s.deploymentCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
	}

	util.ResponseSuccess(c, results, "deployment")
}