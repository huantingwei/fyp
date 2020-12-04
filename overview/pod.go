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

func (s *Service) GetPodInfo(c *gin.Context){
	podInfo := initPodArray();

	insertManyResult, err := s.podCollection.InsertMany(context.TODO(),podInfo);
	if err != nil {
		fmt.Printf(err.Error());
	}else{
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs);
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"type": "pod",
		"data": podInfo,
		"count": len(podInfo),
	});
}

func initPodArray() []interface{}{
	clientset := util.GetKubeClientSet();

	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var podSlice []interface{};

	for _, p := range podList.Items{
		pod := object.Pod{
			ObjectMeta: object.ObjectMeta{
				Name: p.Name,
				Namespace: string(p.Namespace),
				Uid: string(p.UID),
				CreationTime: p.CreationTimestamp.String(),
			},

			DnsPolicy: string(p.Spec.DNSPolicy),
			RestartPolicy: string(p.Spec.RestartPolicy),
			NodeName: p.Spec.NodeName,
			HostIP: p.Status.HostIP,
			PodIP: p.Status.PodIP,
			Phase: string(p.Status.Phase), 
		}

		labelsMap := make(map[string]string);
		var ownerRefSlice []object.OwnerReference;
		var containersSlice []object.Container;
		var containerStatusSlice []object.ContainerStatus;

		for key, val := range p.Labels{
			labelsMap[key] = val;
		}

		for _, o := range p.OwnerReferences{
			owner := object.OwnerReference{
				Name: o.Name,
				Uid: string(o.UID),
				Kind: o.Kind,
			}
			ownerRefSlice = append(ownerRefSlice, owner);
		}

		for _, c := range p.Spec.Containers{
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

		for _, s := range p.Status.ContainerStatuses{
			status := object.ContainerStatus{
				Name: s.Name,
				Image: s.Image,
				Ready: s.Ready,
				RestartCount: int(s.RestartCount),
			}

			containerStatusSlice = append(containerStatusSlice, status);
		}

		pod.Labels = labelsMap;
		pod.OwnerReferences = ownerRefSlice;
		pod.Containers = containersSlice;
		pod.ContainerStatuses = containerStatusSlice;

		podSlice = append(podSlice, pod);
	}
	
	return podSlice;
}