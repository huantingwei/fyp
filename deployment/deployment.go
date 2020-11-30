package deployment

import (
	"fmt"
	"flag"
	"context"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	//appv1 "k8s.io/api/apps/v1"
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
		fmt.Printf("--------Deployment %d: %s--------\n", i, deployment.Name);
		for _, container := range deployment.Spec.Template.Spec.Containers{
			fmt.Printf("Name: %s ; Image: %s\n", container.Name, container.Image);
		}
		fmt.Printf("\n");
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}