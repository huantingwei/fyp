package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/huantingwei/fyp/util"
	"github.com/huantingwei/fyp/object"
)

func (s *Service) refreshReplicaSetInfo(c *gin.Context){
	replicaSetInfo := s.initReplicaSetArray();

	_, err := s.replicaSetCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
        return
	}

	_, err2 := s.replicaSetCollection.InsertMany(context.TODO(),replicaSetInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
        return
	}

	fmt.Println("refreshed replicaSet info")
}

func (s *Service) initReplicaSetArray() []interface{}{
	replicaSetList, err := s.clientset.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var replicaSets []interface{};
	for _, r := range replicaSetList.Items {
		replicaSet := object.ReplicaSet {
			ObjectMeta: object.ObjectMeta{
				Name: r.Name,
				Namespace: string(r.Namespace),
				Uid: string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
			Replicas: int(*r.Spec.Replicas),
			AvailableReplicas: int(r.Status.AvailableReplicas),
			ReadyReplicas: int(r.Status.ReadyReplicas),
		}
		matchLabels := make(map[string]string)
		mlb := r.Spec.Selector.MatchLabels
		for key, val := range mlb {
			matchLabels[key] = val
		}
		replicaSet.MatchLabels = matchLabels
		replicaSets = append(replicaSets, replicaSet)
	}
	

	return replicaSets;
}

func (s *Service) GetReplicaSetInfo(c *gin.Context) {
	cursor, err := s.replicaSetCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
		return
	}

	util.ResponseSuccess(c, results, "replicaSet")
}