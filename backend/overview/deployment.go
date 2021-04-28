package overview

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) refreshDeploymentInfo() error {
	deploymentInfo := s.initDeploymentArray()

	_, err := s.deploymentCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err2 := s.deploymentCollection.InsertMany(context.TODO(), deploymentInfo)
	if err2 != nil {
		return err2
	}

	fmt.Println("refreshed deployment info")
	return nil
}

func (s *Service) initDeploymentArray() []interface{} {
	deploymentList, err := s.clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var deploymentSlice []interface{}

	for _, d := range deploymentList.Items {
		deployment := object.Deployment{
			ObjectMeta: object.ObjectMeta{
				Name:         d.Name,
				Namespace:    string(d.Namespace),
				Uid:          string(d.UID),
				CreationTime: d.CreationTimestamp.String(),
			},
			DesiredPods:         int(*d.Spec.Replicas),
			DNSPolicy:           string(d.Spec.Template.Spec.DNSPolicy),
			RestartPolicy:       string(d.Spec.Template.Spec.RestartPolicy),
			NodeName:            d.Spec.Template.Spec.NodeName,
			UpdatedReplicas:     int(d.Status.UpdatedReplicas),
			ReadyReplicas:       int(d.Status.ReadyReplicas),
			AvailableReplicas:   int(d.Status.AvailableReplicas),
			UnavailableReplicas: int(d.Status.UnavailableReplicas),
		}

		matchLabelsMap := make(map[string]string)
		var containersSlice []object.Container

		matchLabels := d.Spec.Selector.MatchLabels
		for key, val := range matchLabels {
			matchLabelsMap[key] = val
		}

		containers := d.Spec.Template.Spec.Containers
		for _, c := range containers {
			container := object.Container{
				Name:            c.Name,
				Image:           c.Image,
				ImagePullPolicy: string(c.ImagePullPolicy),
			}

			containerPort := make(map[int]string)

			for _, port := range c.Ports {
				containerPort[int(port.ContainerPort)] = string(port.Protocol)
			}
			container.ContainerPorts = containerPort
			containersSlice = append(containersSlice, container)
		}

		deployment.MatchLabels = matchLabelsMap
		deployment.Containers = containersSlice
		deploymentSlice = append(deploymentSlice, deployment)
	}

	return deploymentSlice
}

func (s *Service) GetDeploymentInfo(c *gin.Context) {
	var results []object.Deployment
	cursor, err := s.deploymentCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.Deployment
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "deployment")
}
