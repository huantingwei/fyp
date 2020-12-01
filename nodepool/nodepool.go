package nodepool

import (
	"context"
	"fmt"

	container "cloud.google.com/go/container/apiv1"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"google.golang.org/api/option"
)

func main() {

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
		nodepools := res.GetNodePools();

		for index, pool := range nodepools{
			fmt.Printf("------------Node pool %d------------\n\n", index+1);
			fmt.Printf("------------Node pool general info------------\n");

			fmt.Printf("Name: %s\n", pool.GetName());
			fmt.Printf("Node version: %s\n", pool.GetVersion());
			fmt.Printf("Location: %s\n", pool.GetLocations()[0]);
			fmt.Printf("Status: %d\n", pool.GetStatus());
			fmt.Printf("Autoscaling: %t\n", pool.GetAutoscaling().GetEnabled());
			fmt.Printf("Initial node count: %d\n", pool.GetInitialNodeCount());

			fmt.Printf("------------Node pool configuration------------\n");

			fmt.Printf("Image type: %s\n", pool.GetConfig().GetImageType());
			fmt.Printf("Machine type: %s\n", pool.GetConfig().GetMachineType());
			fmt.Printf("Disk type: %s\n", pool.GetConfig().GetDiskType());
			fmt.Printf("Disk size: %d GB\n", pool.GetConfig().GetDiskSizeGb());

			fmt.Printf("------------Node pool management------------\n");
			fmt.Printf("AutoUpgrade: %t\n", pool.GetManagement().GetAutoUpgrade());
			fmt.Printf("AutoRepair: %t\n", pool.GetManagement().GetAutoRepair());

			fmt.Printf("------------Node pool security------------\n");

			fmt.Printf("Service account: %s\n", pool.GetConfig().GetServiceAccount());
			fmt.Printf("Secure boot enabled: %t\n", pool.GetConfig().GetShieldedInstanceConfig().GetEnableSecureBoot());
			fmt.Printf("Integrity monitoring enabled: %t\n", pool.GetConfig().GetShieldedInstanceConfig().GetEnableIntegrityMonitoring());

		}

	}
}
//fmt.Printf(": %s\n", pool.Get);
