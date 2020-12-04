package overview

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

	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	for index, pod := range podList.Items{
		fmt.Printf("\n---------------------Pod %d---------------------\n", index+1);
		fmt.Printf("\n### Object metadata ###\n");

		fmt.Printf("Node name: %s\n", pod.Name);
		fmt.Printf("Namespace: %s\n", pod.Namespace);
		fmt.Printf("UID: %s\n", pod.UID);
		fmt.Printf("Creation time: %s\n", pod.CreationTimestamp.String());

		fmt.Printf("\n### Object metadata - extra ###\n");
		fmt.Printf("Labels: ");
		for key, val := range pod.Labels{
			fmt.Printf("[%s: %s] ", key, val);
		}
		fmt.Printf("\n");

		fmt.Printf("Owner reference: \n");
		for _, owner := range pod.OwnerReferences{
			fmt.Printf("[Name: %s]\n", owner.Name);
			fmt.Printf("[UID: %s]\n", owner.UID);
			fmt.Printf("[Kind: %s]\n", owner.Kind);
		}
		fmt.Printf("\n");

		fmt.Printf("\n### Pod spec ###\n");
		fmt.Printf("---Containers in the pod---\n");
		containers := pod.Spec.Containers;
		for _, container := range containers{
			fmt.Printf("Container name: %s\n", container.Name);
			fmt.Printf("Container image: %s\n", container.Image);
			fmt.Printf("Container pull policy: %s\n", container.ImagePullPolicy);
			fmt.Printf("Container ports: ");
			for _, port := range container.Ports{
				fmt.Printf("[containerPort: %d ; Protocol: %s] ", port.ContainerPort, port.Protocol);
			}
			fmt.Printf("\n---------------\n");
		}
		fmt.Printf("DNS policy: %s\n", pod.Spec.DNSPolicy);
		fmt.Printf("Restart policy: %s\n", pod.Spec.RestartPolicy);
		fmt.Printf("Node name: %s\n", pod.Spec.NodeName);

		fmt.Printf("\n### Pod status ###\n");
		fmt.Printf("Host IP: %s\n", pod.Status.HostIP);
		fmt.Printf("Pod IP: %s\n", pod.Status.PodIP);
		fmt.Printf("Phase: %s\n", pod.Status.Phase);
		fmt.Printf("Container status: \n");
		for i, container := range pod.Status.ContainerStatuses{
			fmt.Printf("--- Containers %d ---\n", i+1);
			fmt.Printf("Container name: %s\n", container.Name);
			fmt.Printf("Container Image: %s\n", container.Image);
			fmt.Printf("Ready: %t\n", container.Ready);
			fmt.Printf("Container restart count: %d\n", container.RestartCount);

		}

		
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}