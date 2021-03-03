package util

import (
	//standard lib
	"context"
	"fmt"

	//GCP client lib
	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

// const cred = "./util/serviceAccount_ting.json"
// const cluster = `projects/justbadcodes-root/locations/asia-east1-a/clusters/test`
const cluster = `projects/fyp-demo-306511/locations/us-central1-a/clusters/demo`
const cred = "./fyp-demo-sa.json"

func GetGCPClusterManagementClient() *containerpb.Cluster {
	ctx := context.TODO()
	client, err1 := container.NewClusterManagerClient(ctx, option.WithCredentialsFile(cred))

	if err1 != nil {
		fmt.Printf("Fail to create cluster manager client\n")
		fmt.Printf(err1.Error())
	}

	req := &containerpb.GetClusterRequest{
		Name: cluster,
	}

	fmt.Printf("Successfully fetched cluster!!!\n")
	cluster, err2 := (*client).GetCluster(ctx, req)

	if err2 != nil {
		panic(err2.Error())
	}

	return cluster
}
