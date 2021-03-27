package network

import (
	"fmt"
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
	Ports		[]SVPort
}

type Node struct {
	ID			string
	Name		string
	ExternalIP	string
	ClusterIP	string
}

type SVPort struct {
	Port       string 
	NodePort   string
	TargetPort string
	Protocol   string
}

type GraphNode struct {
	ID		string	`json:"id"`
	Kind	string	`json:"kind"`
	Type 	string	`json:"type"`
	Name	string	`json:"name"`
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
			ID: string(n.UID),
			Name: n.Name,
			ClusterIP: n.Spec.PodCIDR,
		}
		for _, a := range n.Status.Addresses {
			if a.Type == "ExternalIP" {
				node.ExternalIP = a.Address
			}
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
	nodes := s.getNodes()
	if nodes == nil {
		panic(fmt.Errorf("No Nodes"))
	}

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
			// no nodePort
			var ports []SVPort
			for _, p := range n.Spec.Ports {
				port := SVPort{
					Port: fmt.Sprintf("%v", p.Port),
					TargetPort: fmt.Sprintf("%v", p.TargetPort.IntVal),
					Protocol: fmt.Sprintf("%v", p.Protocol),
				}
				ports = append(ports, port)
			}
			sv.Ports = ports
			break
		case "LoadBalancer":
			sv.Type = "LoadBalancer"
			for i, ing := range n.Status.LoadBalancer.Ingress {
				if i != 0 {
					sv.ExternalIP += ","
				}
				sv.ExternalIP += ing.IP
			}
			var ports []SVPort
			for _, p := range n.Spec.Ports {
				port := SVPort{
					NodePort: fmt.Sprintf("%v", p.NodePort),
					Port: fmt.Sprintf("%v", p.Port),
					TargetPort: fmt.Sprintf("%v", p.TargetPort.IntVal),
					Protocol: fmt.Sprintf("%v", p.Protocol),
				}
				ports = append(ports, port)
			}
			sv.Ports = ports
			break
		case "NodePort":
			sv.Type = "NodePort"
			// external IP = Node's external IP
			sv.ExternalIP = nodes[0].ExternalIP

			// 
			var ports []SVPort
			for _, p := range n.Spec.Ports {
				port := SVPort{
					NodePort: fmt.Sprintf("%v", p.NodePort),
					Port: fmt.Sprintf("%v", p.Port),
					TargetPort: fmt.Sprintf("%v", p.TargetPort.IntVal),
					Protocol: fmt.Sprintf("%v", p.Protocol),
				}
				ports = append(ports, port)
			}
			sv.Ports = ports
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
	nodes := s.getNodes()
	services := s.getServices(namespace)
	pods := s.getPods(namespace)

	var nodePorts []string

	// client
	clientNode := GraphNode{
		ID:	"ex-client",
		Name: "External Client",
		Type: "Client",
		Kind: "Client",
		Content: "",
	}
	gnodes = append(gnodes, clientNode)

	// Service & Pod
	count := 1
	for _, sv := range services {

		// get nodePorts for building NodeNode later
		if sv.Type == "NodePort"{
			for _, p := range sv.Ports{
				nodePorts = append(nodePorts, p.NodePort)
			}
		}

		// Client -> Service.LoadBalancer
		if sv.Type == "LoadBalancer" {
			for _, port := range sv.Ports {
				link := GraphLink {
					ID: "ex-client" + ":" + sv.ID,
					Source: "ex-client",
					Target: sv.ID,
					Content: sv.ExternalIP + ":" + port.Port,
				}
				glinks = append(glinks, link)
			}
		}

		// Service GraphNode
		svNode := GraphNode{
			ID:	sv.ID,
			Name: sv.Name,
			Kind: "Service",
			Type: sv.Type,
			Content: sv.ClusterIP,
		}
		gnodes = append(gnodes, svNode)

		// Service -> Pod
		for _, pod := range pods {
			for pk, pv := range pod.Labels {
				svv, ok := sv.Selector[pk]
				// Service.Selector == Pod.Labels
				if ok == true && svv == pv {
					// Service.Ports
					for _, port := range sv.Ports{
						link := GraphLink {
							ID: sv.ID + ":" + pod.ID,
							Source: sv.ID,
							Target: pod.ID,
							// port : targetPort
							Content: port.Port + ":" + port.TargetPort,
						}
						glinks = append(glinks, link)
					}
				}
				
			}
			// Pod GraphNode
			if count == 1 {
				podNode := GraphNode {
					ID:	pod.ID,
					Name: pod.Name,
					Type: "Pod",
					Kind: "Pod",
					Content: pod.IP,
				}
				gnodes = append(gnodes, podNode)
			}
			
		}
		count++
	}

	// Node GraphNode
	for _, n := range nodes {
		nodeNode := GraphNode{
			ID:	n.ID,
			Name: n.Name,
			Type: "Node",
			Kind: "Node",
			Content: n.ExternalIP,
		}
		gnodes = append(gnodes, nodeNode)
		// Client -> Node
		// use nodePorts
		for _, np := range nodePorts {
			link := GraphLink {
				ID: "ex-client" + ":" + nodeNode.ID,
				Source: "ex-client",
				Target: nodeNode.ID,
				Content: n.ExternalIP + ":" + np,
			}
			glinks = append(glinks, link)
		}

		// Node -> Service-NodePort
		for _, sv := range services {
			if sv.Type == "NodePort" {
				for _, port := range sv.Ports {
					link := GraphLink {
						ID: n.ID + ":" + sv.ID,
						Source: n.ID,
						Target: sv.ID,
						Content: port.NodePort + ":" + port.Port,
					}
					glinks = append(glinks, link)
				}
			}
		}
	}
	return
}


func (s *Service) GetGraph(c *gin.Context) {
	namespace := c.Query("namespace")

	nodes, links := s.buildGraph(namespace)
	// check null
	if nodes == nil || links == nil {
		err := fmt.Errorf("No Graph available")
		util.ResponseError(c, err)
		return
	}

	// remove unlinked nodes
	for _, node := range nodes {
		linked := 0
		for _, link := range links {
			if link.Source == node.ID || link.Target == node.ID {
				linked = 1
				break
			}
		}
		// not linked to anyone -> remove from nodes
		if linked == 0 {
			//TODO
		}
	}

	graph := Graph{
		Nodes: nodes,
		Links: links,
	}
	util.ResponseSuccess(c, graph, "graph")
}

func (s *Service) GetNamespace(c *gin.Context) {
	namespaces := s.getNamespace()

	if namespaces == nil {
		err := fmt.Errorf("No Namespace available")
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, namespaces, "namepsace")
}
 
