package util

import (
	//standard lib
	"context"
	"fmt"
	"os"
	"bufio"

	//GCP client lib
	container "cloud.google.com/go/container/apiv1"
	"google.golang.org/api/option"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

const clientFile = "./client.txt"

func GetGCPClusterManagementClient() *containerpb.Cluster {
	var err error
	var cred string
	var clusterMeta string

	if f, err := os.Stat(clientFile); err == nil {
		if f.Size() > 0 {
			file, err := os.Open(clientFile)
			if err != nil {
				fmt.Printf("Fail to read clientFile\n")
				panic(err.Error())
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Scan()
			cred = scanner.Text()
			scanner.Scan()
			clusterMeta = scanner.Text()
		}
	} else {
		fmt.Printf("Fail to stat clientFile\n")
		panic(err.Error())
	}

	fmt.Println("cred=", cred)

	ctx := context.TODO()
	client, err := container.NewClusterManagerClient(ctx, option.WithCredentialsFile(cred))

	if err != nil {
		fmt.Printf("Fail to create cluster manager client\n")
		panic(err.Error())
	}

	req := &containerpb.GetClusterRequest{
		Name: clusterMeta,
	}

	
	cluster, err := (*client).GetCluster(ctx, req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Successfully fetched cluster!!!\n")
	
	return cluster
}
