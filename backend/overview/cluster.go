package overview

import (
	"fmt"
	"context"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshClusterInfo(c *gin.Context) {
	clusterInfo := initClusterStruct();

	_, err := s.clusterCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
        return
	}

	_, err2 := s.clusterCollection.InsertOne(context.TODO(),clusterInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
        return
	}
	fmt.Println("refreshed cluster info")
}

func initClusterStruct() object.Cluster{
	client := util.GetGCPClusterManagementClient();

	newCluster := object.Cluster{
	// Get cluster general info
	Name: client.GetName(),
	CreationTime: client.GetCreateTime(),
	MasterVersion: client.GetCurrentMasterVersion(),
	IPendpoint: client.GetEndpoint(),
	Location: client.GetLocation(),
	ReleaseChannel: int(client.GetReleaseChannel().GetChannel()),
	Status: client.GetStatus().String(),

	//Get cluster networking config
	Network: client.GetNetwork(),
	NetworkConfig: client.GetNetworkConfig().GetNetwork(),
	Subnet: client.GetNetworkConfig().GetSubnetwork(),
	IntranodeVisibility: client.GetNetworkConfig().GetEnableIntraNodeVisibility(),
	NetworkPolicyEnabled: client.GetNetworkPolicy().GetEnabled(),
	MasterAuthNetworkEnabled: client.GetMasterAuthorizedNetworksConfig().GetEnabled(),

	//Get cluster security config
	ShieldedNodeEnabled: client.GetShieldedNodes().GetEnabled(),
	BinaryAuthorisationEnabled: client.GetBinaryAuthorization().GetEnabled(),
	ClientCertificateEnabled: client.GetMasterAuth().GetClientCertificateConfig().GetIssueClientCertificate(),
	}
	
	return newCluster;
}

func (s *Service) GetClusterInfo(c *gin.Context) {
	cursor, err := s.clusterCollection.Find(context.TODO(), bson.D{})
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

	util.ResponseSuccess(c, results, "cluster")
}