package network

import (
	"fmt"
	"errors"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
)

type Pod struct {
	ID			string
	Name		string
	IP			string
	Labels 		map[string]string
}

type SV struct {
	ID			string
	Name		string		
	Type		string		
	ClusterIP	string		
	ExternalIP	string		
	Selector	map[string]string
	//Ports	[]SVPort	`json:"ports"`	
}

type Node struct {
	Name		string
	IP			string
}

//type SVPort struct {
//	Port       int    `json:"port"`
//	NodePort   int    `json:"nodePort"`
//	TargetPort int    `json:"targetPort"`
//	Protocol   string `json:"protocol"`
//}

type GraphNode struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
	Type 	string	`json:"type"`
	Content	string	`json:"content"`
}

type GraphLink struct {
	ID		string	`json:"id"`
	Source	string	`json:"source"`
	Target 	string	`json:"target"`
	Content	string	`json:"content"`
}

type Graph struct {
	Nodes	[]GraphNode	`json:"nodes"`
	Links	[]GraphLink	`json:"links"`
}

func (s *Service) getNodes() (nodes []Node){
	nodeList, err := s.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}
	for _, n := range nodeList.Items {
		node := Node {
			Name: n.Name,
			IP: n.Spec.PodCIDR,
		}
		nodes = append(nodes, node)
	}
	return
}

func (s *Service) getNamespace() (namespaces []string) {
	npList, err := s.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}
	for _, n := range npList.Items {
		namespaces = append(namespaces, n.ObjectMeta.Name)
	}
	return
}

func (s *Service) getServices(namespace string) (services []SV) {
	serviceList, err := s.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	for _, n := range serviceList.Items {
		sv := SV {
			ID: string(n.UID),
			Name: n.Name,
			ClusterIP: n.Spec.ClusterIP,
		}
		selectors := make(map[string]string)
		for key, val := range n.Spec.Selector{
			selectors[key] = val;
		}
		sv.Selector = selectors

		switch t := string(n.Spec.Type); t{
		case "ClusterIP":
			sv.Type =  "ClusterIP"
			sv.ExternalIP = ""
			break
		case "LoadBalancer":
			sv.Type = "LoadBalancer"
			for i, ing := range n.Status.LoadBalancer.Ingress {
				if i != 0 {
					sv.ExternalIP += ","
				}
				sv.ExternalIP += ing.IP
			}
			break
		case "NodePort":
			sv.Type = "NodePort"
			sv.ExternalIP = ""
			break
		}
		services = append(services, sv)
	}
	return 
}

func (s *Service) getPods(namespace string) (pods []Pod){
	podList, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	for _, n := range podList.Items{
		pod := Pod {
			ID: string(n.UID),
			Name: n.Name,
			IP: n.Status.PodIP,
		}
		labels := make(map[string]string)
		for key, val := range n.Labels{
			labels[key] = val;
		}
		pod.Labels = labels
		pods = append(pods, pod)
	}

	return pods
}

func (s *Service) buildGraph(namespace string) (gnodes []GraphNode, glinks []GraphLink){
	services := s.getServices(namespace)
	pods := s.getPods(namespace)

	count := 1
	for _, sv := range services {

		svNode := GraphNode {
			ID:	sv.ID,
			Name: sv.Name,
			Type: "Service",
			Content: "ClusterIP: " + sv.ClusterIP,
		}
		if sv.ExternalIP != "" {
			svNode.Content += (" / ExternalIP: " + sv.ExternalIP)
		}
		gnodes = append(gnodes, svNode)

		// Service -- Pod
		for _, pod := range pods {
			for pk, pv := range pod.Labels {
				svv, ok := sv.Selector[pk]
				if ok == true && svv == pv {
					link := GraphLink {
						ID: sv.ID + ":" + pod.ID,
						Source: sv.ID,
						Target: pod.ID,
						Content: pk + ":" + pv,
					}
					glinks = append(glinks, link)
				}
				
			}
			// Pod GraphNode
			if count == 1 {
				podNode := GraphNode {
					ID:	pod.ID,
					Name: pod.Name,
					Type: "Pod",
					Content: pod.IP,
				}
				gnodes = append(gnodes, podNode)
			}
			
		}
		count++
	}

	return
}

func (s *Service) GetGraph(c *gin.Context) {
	namespace := c.Query("namespace")

	nodes, links := s.buildGraph(namespace)
	// check null
	if nodes == nil || links == nil {
		err := errors.New("No Graph available")
		util.ResponseError(c, err)
		return
	}

	graph := Graph{
		Nodes: nodes,
		Links: links,
	}
	fmt.Println("hi")
	util.ResponseSuccess(c, graph, "graph")
}

func (s *Service) GetNamespace(c *gin.Context) {
	namespaces := s.getNamespace()
	
	if namespaces == nil {
		err := errors.New("No Namespace available")
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, namespaces, "namepsace")
}
 
