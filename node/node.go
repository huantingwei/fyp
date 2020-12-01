package node

import (
	"fmt"
	"flag"
	"context"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

)

func main(){
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	nodeList, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	for index, node := range nodeList.Items{
		fmt.Printf("\n---------------------Node %d---------------------\n", index+1);
		fmt.Printf("\n### Object metadata ###\n");

		fmt.Printf("Node name: %s\n", node.Name);
		fmt.Printf("Namespace: %s\n", node.Namespace);
		fmt.Printf("UID: %s\n", node.UID);
		fmt.Printf("Creation time: %s\n", node.CreationTimestamp.String());

		fmt.Printf("\n### Node Spec ###\n");
		fmt.Printf("Pod CIDR: %s\n", node.Spec.PodCIDR);
		fmt.Printf("Node ID: %s\n", node.Spec.ProviderID);

		fmt.Printf("\n### Node System info ###\n");
		fmt.Printf("Machine ID: %s\n", node.Status.NodeInfo.MachineID);
		fmt.Printf("Kernel version: %s\n", node.Status.NodeInfo.KernelVersion);
		fmt.Printf("OS Image: %s\n", node.Status.NodeInfo.OSImage);
		fmt.Printf("Operating system: %s\n", node.Status.NodeInfo.OperatingSystem);
		fmt.Printf("Container runtime version: %s\n", node.Status.NodeInfo.ContainerRuntimeVersion);
		fmt.Printf("Kubelet Version: %s\n", node.Status.NodeInfo.KubeletVersion);
		fmt.Printf("Kube-proxy Version: %s\n", node.Status.NodeInfo.KubeProxyVersion);

		fmt.Printf("\n### Node resources capacity ###\n");
		fmt.Printf("CPU: %d CPU\n", node.Status.Capacity.Cpu().Value());
		fmt.Printf("Memory: %.2f GB\n", float64(node.Status.Capacity.Memory().Value())/1000000000.0);
		fmt.Printf("Pod: %d pods\n", node.Status.Capacity.Pods().Value());
		fmt.Printf("Ephemeral storage: %.2f GB\n", float64(node.Status.Capacity.StorageEphemeral().Value())/1000000000.0);
		fmt.Printf("Storage: %d B\n", node.Status.Capacity.Storage().Value());

		fmt.Printf("\n### Node resources allocatable ###\n");
		fmt.Printf("CPU: %d CPU\n", node.Status.Allocatable.Cpu().Value());
		fmt.Printf("Memory: %.2f GB\n", float64(node.Status.Allocatable.Memory().Value())/1000000000.0);
		fmt.Printf("Pod: %d pods\n", node.Status.Allocatable.Pods().Value());
		fmt.Printf("Ephemeral storage: %.2f GB\n", float64(node.Status.Allocatable.StorageEphemeral().Value())/1000000000.0);
		fmt.Printf("Storage: %d B\n", node.Status.Allocatable.Storage().Value());

		fmt.Printf("\n### Node conditions ###\n");
		for index, condition := range node.Status.Conditions{
			fmt.Printf("Condition %d \n Condition name: %s \n Status: %s \n Last heartbeat: %s \n Last transition: %s \n Message: %s\n\n",
			index, condition.Type, condition.Status, condition.LastHeartbeatTime.String(), condition.LastTransitionTime.String(), condition.Message)
		}
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}