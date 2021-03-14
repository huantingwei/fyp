package util

import (
	//standard lib
	"context"
	"fmt"
	"os"

	//GCP client lib
	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func GetGCPClusterManagementClient() *containerpb.Cluster {
	clusterMeta := os.Getenv("CLUSTER")
	cred := os.Getenv("CRED")

	ctx := context.TODO()
	client, err1 := container.NewClusterManagerClient(ctx, option.WithCredentialsFile(cred))

	if err1 != nil {
		fmt.Printf("Fail to create cluster manager client\n")
		fmt.Printf(err1.Error())
	}

	req := &containerpb.GetClusterRequest{
		Name: clusterMeta,
	}

	fmt.Printf("Successfully fetched cluster!!!\n")
	cluster, err2 := (*client).GetCluster(ctx, req)

	if err2 != nil {
		panic(err2.Error())
	}

	return cluster
}
