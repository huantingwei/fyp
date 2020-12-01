package deployment

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

	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	for i, deployment := range deploymentList.Items{
		fmt.Printf("\n-------------Deployment %d------------\n", i+1);
		fmt.Printf("\n### Deployment metadata ###\n");
		fmt.Printf("Name: %s\n", deployment.Name);
		fmt.Printf("Namespace: %s\n", deployment.Namespace);
		fmt.Printf("UID: %s\n", deployment.UID);
		fmt.Printf("Creation time: %s\n", deployment.CreationTimestamp.String());

		fmt.Printf("\n### Deployment spec ###\n");
		fmt.Printf("Number of desired pods: %d\n", deployment.Spec.Replicas);
		fmt.Printf("Label selector - match labels: ");
		matchLabels := deployment.Spec.Selector.MatchLabels;
		for key, val := range matchLabels{
			fmt.Printf("[%s: %s] ", key, val);
		}
		fmt.Printf("\n");

		fmt.Printf("\n### Pod specification of this deployment ###\n");
		fmt.Printf("---Containers in the pod---\n");
		containers := deployment.Spec.Template.Spec.Containers;
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
		fmt.Printf("DNS policy: %s\n", deployment.Spec.Template.Spec.DNSPolicy);
		fmt.Printf("Restart policy: %s\n", deployment.Spec.Template.Spec.RestartPolicy);

		fmt.Printf("\n### Deployment status ###\n");
		fmt.Printf("Updated replicas: %d\n", deployment.Status.UpdatedReplicas);
		fmt.Printf("Ready replicas: %d\n", deployment.Status.ReadyReplicas);
		fmt.Printf("Available replicas: %d\n", deployment.Status.AvailableReplicas);
		fmt.Printf("Unavailable replicas: %d\n", deployment.Status.UnavailableReplicas);

	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}