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

