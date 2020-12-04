package overview

import (
	//standard lib
	"fmt"
	"context"
	"net/http"
	//GCP client lib
	//containerpb "google.golang.org/genproto/googleapis/container/v1"
	//internal package
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	//Gin
	"github.com/gin-gonic/gin"
)

func (s *Service) GetNodepoolInfo(c *gin.Context) {
	nodepoolInfo := initNodepoolsArray();

	insertManyResult, err := s.nodepoolCollection.InsertMany(context.TODO(),nodepoolInfo);
	if err != nil {
		fmt.Printf(err.Error());
	}else{
		fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs);
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"type": "nodepool",
		"data": nodepoolInfo,
		"count": len(nodepoolInfo),
	});
}

func initNodepoolsArray() []interface{}{
	client := util.GetGCPClusterManagementClient();
	nodepools := client.GetNodePools();
	var nodepoolSlice []interface{};

	for _, pool := range nodepools{

		newNodepool := object.Nodepool{
			Name: pool.GetName(),
			Version: pool.GetVersion(),
			Location: pool.GetLocations()[0],
			Status: int(pool.GetStatus()),
			AutoscalingEnabled: pool.GetAutoscaling().GetEnabled(),
			InitialNodeCount: int(pool.GetInitialNodeCount()),
	
			//Node pool configuration
			ImageType: pool.GetConfig().GetImageType(),
			MachineType: pool.GetConfig().GetMachineType(),
			DiskType: pool.GetConfig().GetDiskType(),
			DiskSize: int(pool.GetConfig().GetDiskSizeGb()),
	
			//Node pool management
			AutoUpgrade: pool.GetManagement().GetAutoUpgrade(),
			AutoRepair: pool.GetManagement().GetAutoRepair(),
	
			//Node pool security
			ServiceAccount: pool.GetConfig().GetServiceAccount(),
			SecureBoot: pool.GetConfig().GetShieldedInstanceConfig().GetEnableSecureBoot(),
			IntegrityMonitoring: pool.GetConfig().GetShieldedInstanceConfig().GetEnableIntegrityMonitoring(),
	
		}
		nodepoolSlice = append(nodepoolSlice, newNodepool);
	}
	
	return nodepoolSlice;
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