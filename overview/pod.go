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

func (s *Service) refreshPodInfo(c *gin.Context){
	podInfo := s.initPodArray();

	_, err := s.podCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	_, err2 := s.podCollection.InsertMany(context.TODO(),podInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
	}

	fmt.Println("refreshed pod info")
}

func (s *Service) initPodArray() []interface{}{
	// clientset := util.GetKubeClientSet();

	podList, err := s.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{});
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

			DNSPolicy: string(p.Spec.DNSPolicy),
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

func (s *Service) GetPodInfo(c *gin.Context) {
	cursor, err := s.podCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
	}

	util.ResponseSuccess(c, results, "pod")
}