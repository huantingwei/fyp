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

func (s *Service) refreshServiceInfo() error {
	serviceInfo := s.initServiceArray()

	_, err := s.serviceCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err2 := s.serviceCollection.InsertMany(context.TODO(), serviceInfo)
	if err2 != nil {
		return err2
	}

	fmt.Println("refreshed service info")
	return nil
}

func (s *Service) initServiceArray() []interface{} {
	//clientset := util.GetKubeClientSet();

	serviceList, err := s.clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var serviceSlice []interface{}

	for _, s := range serviceList.Items {
		service := object.Service{
			ObjectMeta: object.ObjectMeta{
				Name:         s.Name,
				Namespace:    string(s.Namespace),
				Uid:          string(s.UID),
				CreationTime: s.CreationTimestamp.String(),
			},
			ClusterIP:   s.Spec.ClusterIP,
			ServiceType: string(s.Spec.Type),
		}

		labelMap := make(map[string]string)
		var servicePortsSlice []object.ServicePort
		var ingressIPSlice []string

		for key, val := range s.Spec.Selector {
			labelMap[key] = val
		}

		for _, p := range s.Spec.Ports {
			port := object.ServicePort{
				Port:       int(p.Port),
				NodePort:   int(p.NodePort),
				TargetPort: int(p.TargetPort.IntVal),
				Protocol:   string(p.Protocol),
			}
			servicePortsSlice = append(servicePortsSlice, port)
		}

		for _, ingress := range s.Status.LoadBalancer.Ingress {
			ingressIPSlice = append(ingressIPSlice, ingress.IP)
		}

		service.LabelSelectors = labelMap
		service.ServicePorts = servicePortsSlice
		service.IngressIP = ingressIPSlice

		serviceSlice = append(serviceSlice, service)
	}
	return serviceSlice
}

func (s *Service) GetServiceInfo(c *gin.Context) {
	var results []object.Service
	var tmp object.Service
	cursor, err := s.serviceCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "service")
}
