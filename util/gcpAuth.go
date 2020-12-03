package util

import (
	//standard lib
	"context"
	"fmt"

	//GCP client lib
	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"google.golang.org/api/option"
)

func GetGCPClusterManagementClient() *containerpb.Cluster{
	ctx := context.TODO();
	client, err1 := container.NewClusterManagerClient(ctx, option.WithCredentialsFile("./util/serviceAccount.json"));

	if err1 != nil {
		fmt.Printf("Fail to create cluster manager client\n")
		fmt.Printf(err1.Error());
	}

	req := &containerpb.GetClusterRequest{
		Name: `projects/fyp-gcp-296605/locations/us-central1-c/clusters/cluster-1`,
	}

	cluster, err2 := (*client).GetCluster(ctx, req)

	if err2 != nil {
		panic(err2.Error())
	}

	return cluster;	
}