package overview

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshClusterInfo() error {
	clusterInfo := initClusterStruct()

	_, err := s.clusterCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err2 := s.clusterCollection.InsertOne(context.TODO(), clusterInfo)
	if err2 != nil {
		return err2
	}
	fmt.Println("refreshed cluster info")
	return nil
}

func initClusterStruct() object.Cluster {
	client := util.GetGCPClusterManagementClient()

	newCluster := object.Cluster{
		// Get cluster general info
		Name:           client.GetName(),
		CreationTime:   client.GetCreateTime(),
		MasterVersion:  client.GetCurrentMasterVersion(),
		IPendpoint:     client.GetEndpoint(),
		Location:       client.GetLocation(),
		ReleaseChannel: int(client.GetReleaseChannel().GetChannel()),
		Status:         client.GetStatus().String(),

		//Get cluster networking config
		Network:                  client.GetNetwork(),
		NetworkConfig:            client.GetNetworkConfig().GetNetwork(),
		Subnet:                   client.GetNetworkConfig().GetSubnetwork(),
		IntranodeVisibility:      client.GetNetworkConfig().GetEnableIntraNodeVisibility(),
		NetworkPolicyEnabled:     client.GetNetworkPolicy().GetEnabled(),
		MasterAuthNetworkEnabled: client.GetMasterAuthorizedNetworksConfig().GetEnabled(),

		//Get cluster security config
		ShieldedNodeEnabled:        client.GetShieldedNodes().GetEnabled(),
		BinaryAuthorisationEnabled: client.GetBinaryAuthorization().GetEnabled(),
		ClientCertificateEnabled:   client.GetMasterAuth().GetClientCertificateConfig().GetIssueClientCertificate(),
	}

	return newCluster
}

func (s *Service) GetClusterInfo(c *gin.Context) {
	var result object.Cluster
	err := s.clusterCollection.FindOne(context.TODO(), bson.M{}).Decode(&result)
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, result, "cluster")
}
