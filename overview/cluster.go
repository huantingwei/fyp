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

func (s *Service) GetClusterInfo(c *gin.Context) {
	gcpClient := util.GetGCPClusterManagementClient();
	newCluster := initClusterStruct(gcpClient);
	printCluster(newCluster);

	insertion, err := s.clusterCollection.InsertOne(context.TODO(),newCluster);
	if err != nil {
		fmt.Printf(err.Error());
	}
	
	fmt.Println("Inserted a single document: ", insertion.InsertedID);

	c.IndentedJSON(http.StatusOK, newCluster);
}

func initClusterStruct(client *containerpb.Cluster) object.Cluster{
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

