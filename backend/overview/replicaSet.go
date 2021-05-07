package overview

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) refreshReplicaSetInfo() error {
	replicaSetInfo := s.initReplicaSetArray()

	_, err := s.replicaSetCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err2 := s.replicaSetCollection.InsertMany(context.TODO(), replicaSetInfo)
	if err2 != nil {
		return err2
	}

	fmt.Println("refreshed replicaSet info")
	return nil
}

func (s *Service) initReplicaSetArray() []interface{} {
	replicaSetList, err := s.clientset.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var replicaSets []interface{}
	for _, r := range replicaSetList.Items {
		replicaSet := object.ReplicaSet{
			ObjectMeta: object.ObjectMeta{
				Name:         r.Name,
				Namespace:    string(r.Namespace),
				Uid:          string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
			Replicas:          int(*r.Spec.Replicas),
			AvailableReplicas: int(r.Status.AvailableReplicas),
			ReadyReplicas:     int(r.Status.ReadyReplicas),
		}
		matchLabels := make(map[string]string)
		mlb := r.Spec.Selector.MatchLabels
		for key, val := range mlb {
			matchLabels[key] = val
		}
		replicaSet.MatchLabels = matchLabels
		replicaSets = append(replicaSets, replicaSet)
	}

	return replicaSets
}

func (s *Service) GetReplicaSetInfo(c *gin.Context) {
	var results []object.ReplicaSet
	cursor, err := s.replicaSetCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.ReplicaSet
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "replicaSet")
}
