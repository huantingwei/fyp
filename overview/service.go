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

func (s *Service) GetServiceInfo(c *gin.Context){
	serviceInfo := initServiceArray();

	insertManyResult, err := s.serviceCollection.InsertMany(context.TODO(),serviceInfo);
	if err != nil {
		fmt.Printf(err.Error());
	}else{
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs);
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"type": "service",
		"data": serviceInfo,
		"count": len(serviceInfo),
	});
	
}

func initServiceArray() []interface{}{
	clientset := util.GetKubeClientSet();

	serviceList, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var serviceSlice []interface{};

	for _, s := range serviceList.Items{
		service := object.Service{
			ObjectMeta: object.ObjectMeta{
				Name: s.Name,
				Namespace: string(s.Namespace),
				Uid: string(s.UID),
				CreationTime: s.CreationTimestamp.String(),
			},
			ClusterIP: s.Spec.ClusterIP,
			ServiceType: string(s.Spec.Type),
		}

		labelMap := make(map[string]string);
		var servicePortsSlice []object.ServicePort;
		var ingressIPSlice []string;

		for key, val := range s.Spec.Selector{
			labelMap[key] = val;
		}

		for _, p := range s.Spec.Ports{
			port := object.ServicePort{
				Port: int(p.Port),
				NodePort: int(p.NodePort),
				TargetPort: int(p.TargetPort.IntVal),
				Protocol: string(p.Protocol),
			}
			servicePortsSlice = append(servicePortsSlice, port);
		}	

		for _, ingress := range s.Status.LoadBalancer.Ingress{
			ingressIPSlice = append(ingressIPSlice, ingress.IP)
		}

		service.LabelSelectors = labelMap;
		service.ServicePorts = servicePortsSlice;
		service.IngressIP = ingressIPSlice;

		serviceSlice = append(serviceSlice, service);
	}
	return serviceSlice;
}