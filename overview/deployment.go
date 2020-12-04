package overview

import (
	//standard lib
	"fmt"
	"context"
	"net/http"
	//client lib
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//internal package
	"github.com/huantingwei/fyp/util"
	"github.com/huantingwei/fyp/object"
	//gin
	"github.com/gin-gonic/gin"
)

func (s *Service) GetDeploymentInfo(c *gin.Context){
	deploymentInfo := initDeploymentArray();

	insertManyResult, err := s.deploymentCollection.InsertMany(context.TODO(),deploymentInfo);
	if err != nil {
		fmt.Printf(err.Error());
	}else{
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs);
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"type": "deployment",
		"data": deploymentInfo,
		"count": len(deploymentInfo),
	});
}

func initDeploymentArray() []interface{}{
	clientset := util.GetKubeClientSet();
	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{});
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
