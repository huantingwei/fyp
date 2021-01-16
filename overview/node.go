package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshNodeInfo(c *gin.Context){
	nodeInfo := s.initNodeArray();

	_, err := s.nodeCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	_, err2 := s.nodeCollection.InsertMany(context.TODO(),nodeInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
	}

	fmt.Println("refreshed node info")
}

func (s *Service) initNodeArray() []interface{}{
	//clientset := util.GetKubeClientSet();

	nodeList, err := s.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var nodeSlice []interface{};

	for _, n := range nodeList.Items{
		node := object.Node{
			ObjectMeta: object.ObjectMeta{
				Name: n.Name,
				Namespace: string(n.Namespace),
				Uid: string(n.UID),
				CreationTime: n.CreationTimestamp.String(),
			},
			PodCIDR: n.Spec.PodCIDR,
			NodeID: string(n.Spec.ProviderID),
			MachineID: n.Status.NodeInfo.MachineID,
			KernelVersion: n.Status.NodeInfo.KernelVersion,
			OsImage: n.Status.NodeInfo.OSImage,
			Os: n.Status.NodeInfo.OperatingSystem,
			ContainerRuntime: n.Status.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion: n.Status.NodeInfo.KubeletVersion,
			KubeProxyVersion: n.Status.NodeInfo.KubeProxyVersion,
			CpuCap: int(n.Status.Capacity.Cpu().Value()),
			MemoryCap: float64(n.Status.Capacity.Memory().Value())/1000000000.0,
			PodsCap: int(n.Status.Capacity.Pods().Value()),
			EphemeralStorageCap: float64(n.Status.Capacity.StorageEphemeral().Value())/1000000000.0,
			StorageCap: int(n.Status.Capacity.Storage().Value()),
			CpuAllocatable: int(n.Status.Allocatable.Cpu().Value()),
			MemoryAllocatable: float64(n.Status.Allocatable.Memory().Value())/1000000000.0,
			PodsAllocatable: int(n.Status.Allocatable.Pods().Value()),
			EphemeralStorageAllocatable: float64(n.Status.Allocatable.StorageEphemeral().Value())/1000000000.0,
			StorageAllocatable: int(n.Status.Allocatable.Storage().Value()),
		}

		var conditionSlice []object.Condition;

		for _, c := range n.Status.Conditions{
			condition := object.Condition{
				ConditionName: string(c.Type),
				Status: string(c.Status),
				LastHeartbeatTime: c.LastHeartbeatTime.String(),
				LastTransitionTime: c.LastTransitionTime.String(),
				Message: c.Message,
			}
			conditionSlice = append(conditionSlice, condition);
		}

		node.Conditions = conditionSlice;
		nodeSlice = append(nodeSlice, node);
	}
	return nodeSlice;
}


func (s *Service) GetNodeInfo(c *gin.Context) {
	cursor, err := s.nodeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
	}

	util.ResponseSuccess(c, results, "node")
}