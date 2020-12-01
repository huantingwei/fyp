package cluster
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
		fmt.Printf("------------Cluster general info------------\n");
		fmt.Printf("Name: %s\n", res.GetName());
		fmt.Printf("Creation time: %s\n", res.GetCreateTime());
		fmt.Printf("Version: %s\n", res.GetCurrentMasterVersion());
		fmt.Printf("End point: %s\n", res.GetEndpoint());
		fmt.Printf("Location: %s\n", res.GetLocation());
		fmt.Printf("Release channel: %d\n", res.GetReleaseChannel().GetChannel());
		fmt.Printf("Status: %s\n", res.GetStatus().String());

		fmt.Printf("------------Cluster networking config------------\n");
		fmt.Printf("Network: %s\n", res.GetNetwork());
		fmt.Printf("Network config: %s\n", res.GetNetworkConfig().GetNetwork());
		fmt.Printf("Subnet: %s\n", res.GetNetworkConfig().GetSubnetwork());
		fmt.Printf("Intranode visibility: %t\n", res.GetNetworkConfig().GetEnableIntraNodeVisibility());
		fmt.Printf("Network policy enabled: %t\n", res.GetNetworkPolicy().GetEnabled());
		fmt.Printf("Master authorised network enabled: %t\n", res.GetMasterAuthorizedNetworksConfig().GetEnabled());

		fmt.Printf("------------Cluster security config------------\n");
		fmt.Printf("Shielded node enabled: %t\n", res.GetShieldedNodes().GetEnabled());
		fmt.Printf("Binary authorisation enabled: %t\n", res.GetBinaryAuthorization().GetEnabled());
		fmt.Printf("Client certificate enabled: %t\n", res.GetMasterAuth().GetClientCertificateConfig().GetIssueClientCertificate());


	}
}
