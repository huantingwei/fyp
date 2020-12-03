package overview

import (
	//standard lib
	"fmt"
	"context"
	"net/http"
	//GCP client lib
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	//internal package
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	//Gin
	"github.com/gin-gonic/gin"
)

func (s *Service) GetNodepoolInfo(c *gin.Context) {
	gcpClient := util.GetGCPClusterManagementClient();
	newNodepool := initNodepoolStruct(gcpClient);
	//printCluster(newCluster);

	insertion, err := s.nodepoolCollection.InsertOne(context.TODO(),newNodepool);
	if err != nil {
		fmt.Printf(err.Error());
	}
	
	fmt.Println("Inserted a single document: ", insertion.InsertedID);

	c.IndentedJSON(http.StatusOK, newNodepool);
}

func initNodepoolStruct(client *containerpb.Cluster) object.Nodepool{
	newNodepool := object.Nodepool{
		Name: client.GetName(),
		Version: client.GetVersion(),
		Location: client.GetLocations()[0],
		Status: client.GetStatus(),
		AutoscalingEnabled: client.GetAutoscaling().GetEnabled(),
		InitialNodeCount: client.GetInitialNodeCount(),

		//Node pool configuration
		ImageType: client.GetConfig().GetImageType(),
		MachineType: client.GetConfig().GetMachineType(),
		DiskType: client.GetConfig().GetDiskType(),
		DiskSize: client.GetConfig().GetDiskSizeGb(),

		//Node pool management
		AutoUpgrade: client.GetManagement().GetAutoUpgrade(),
		AutoRepair: client.GetManagement().GetAutoRepair(),

		//Node pool security
		ServiceAccount: client.GetConfig().GetServiceAccount(),
		SecureBoot: client.GetConfig().GetShieldedInstanceConfig().GetEnableSecureBoot(),
		IntegrityMonitoring: client.GetConfig().GetShieldedInstanceConfig().GetEnableIntegrityMonitoring(),

	}

	return newNodepool
}


// nodepools := res.GetNodePools();

// for index, pool := range nodepools{
// 	fmt.Printf("------------Node pool %d------------\n\n", index+1);
// 	fmt.Printf("------------Node pool general info------------\n");

// 	fmt.Printf("Name: %s\n", pool.GetName());
// 	fmt.Printf("Node version: %s\n", pool.GetVersion());
// 	fmt.Printf("Location: %s\n", pool.GetLocations()[0]);
// 	fmt.Printf("Status: %d\n", pool.GetStatus());
// 	fmt.Printf("Autoscaling: %t\n", pool.GetAutoscaling().GetEnabled());
// 	fmt.Printf("Initial node count: %d\n", pool.GetInitialNodeCount());

// 	fmt.Printf("------------Node pool configuration------------\n");

// 	fmt.Printf("Image type: %s\n", pool.GetConfig().GetImageType());
// 	fmt.Printf("Machine type: %s\n", pool.GetConfig().GetMachineType());
// 	fmt.Printf("Disk type: %s\n", pool.GetConfig().GetDiskType());
// 	fmt.Printf("Disk size: %d GB\n", pool.GetConfig().GetDiskSizeGb());

// 	fmt.Printf("------------Node pool management------------\n");
// 	fmt.Printf("AutoUpgrade: %t\n", pool.GetManagement().GetAutoUpgrade());
// 	fmt.Printf("AutoRepair: %t\n", pool.GetManagement().GetAutoRepair());

// 	fmt.Printf("------------Node pool security------------\n");

// 	fmt.Printf("Service account: %s\n", pool.GetConfig().GetServiceAccount());
// 	fmt.Printf("Secure boot enabled: %t\n", pool.GetConfig().GetShieldedInstanceConfig().GetEnableSecureBoot());
// 	fmt.Printf("Integrity monitoring enabled: %t\n", pool.GetConfig().GetShieldedInstanceConfig().GetEnableIntegrityMonitoring());