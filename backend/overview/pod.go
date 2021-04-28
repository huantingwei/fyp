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

func (s *Service) refreshPodInfo() error {
	podInfo := s.initPodArray()

	_, err := s.podCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err2 := s.podCollection.InsertMany(context.TODO(), podInfo)
	if err2 != nil {
		return err2
	}

	fmt.Println("refreshed pod info")
	return nil
}

func (s *Service) initPodArray() []interface{} {
	podList, err := s.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var podSlice []interface{}

	for _, p := range podList.Items {
		pod := object.Pod{
			ObjectMeta: object.ObjectMeta{
				Name:         p.Name,
				Namespace:    string(p.Namespace),
				Uid:          string(p.UID),
				CreationTime: p.CreationTimestamp.String(),
			},

			DNSPolicy:     string(p.Spec.DNSPolicy),
			RestartPolicy: string(p.Spec.RestartPolicy),
			NodeName:      p.Spec.NodeName,
			HostIP:        p.Status.HostIP,
			PodIP:         p.Status.PodIP,
			Phase:         string(p.Status.Phase),
		}

		labelsMap := make(map[string]string)
		var ownerRefSlice []object.OwnerReference
		var controlledBy string
		var containersSlice []object.Container
		var containerStatusSlice []object.ContainerStatus

		for key, val := range p.Labels {
			labelsMap[key] = val
		}

		for i, o := range p.ObjectMeta.OwnerReferences {
			owner := object.OwnerReference{
				Name: o.Name,
				Uid:  string(o.UID),
				Kind: o.Kind,
			}
			ownerRefSlice = append(ownerRefSlice, owner)
			if i != 0 {
				controlledBy += ", "
			}
			controlledBy += o.Name
		}

		for _, c := range p.Spec.Containers {
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

		for _, s := range p.Status.ContainerStatuses {
			status := object.ContainerStatus{
				Name:         s.Name,
				Image:        s.Image,
				Ready:        s.Ready,
				RestartCount: int(s.RestartCount),
			}

			containerStatusSlice = append(containerStatusSlice, status)
		}

		pod.Labels = labelsMap
		pod.OwnerReferences = ownerRefSlice
		pod.ControlledBy = controlledBy
		pod.Containers = containersSlice
		pod.ContainerStatuses = containerStatusSlice

		podSlice = append(podSlice, pod)
	}

	return podSlice
}

func (s *Service) GetPodInfo(c *gin.Context) {
	var results []object.Pod
	cursor, err := s.podCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.Pod
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "pod")
}
