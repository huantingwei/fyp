package overview

import (
	"fmt"
	"context"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshNodepoolInfo() error {
	nodepoolInfo := initNodepoolsArray();

	_, err := s.nodepoolCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
        return err
	}

	_, err2 := s.nodepoolCollection.InsertMany(context.TODO(),nodepoolInfo);
	if err2 != nil {
        return err2
	}

	fmt.Println("refreshed nodepool info")
	return nil
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

func (s *Service) GetNodepoolInfo(c *gin.Context) {
	cursor, err := s.nodepoolCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
		return
	}

	util.ResponseSuccess(c, results, "nodepool")
}
