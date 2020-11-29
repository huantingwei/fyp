package cluster

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"google.golang.org/api/option"
)

func getClusterInfo() {

	ctx := context.TODO();
	client, err1 := container.NewClusterManagerClient(ctx, option.WithCredentialsFile("./sa.json"));

	if err1 != nil {
		fmt.Printf("Fail to create cluster manager client\n")
		fmt.Printf(err1.Error());
	}

	req := &containerpb.GetClusterRequest{
		Name: `projects/fyp-gcp-296605/locations/us-central1-c/clusters/cluster-1`,
	}

	res, err2 := (*client).GetCluster(ctx, req)

	if err2 != nil{
		fmt.Printf("Fail to get cluster\n")
		fmt.Printf(err2.Error());
	}else{
		fmt.Printf("Cluster name: %s\nCluster location: %s\n", (*res).GetName(), (*res).GetLocation());
	}

}
