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

func (s *Service) refreshStatefulSetInfo() error {
	statefulSetInfo := s.initStatefulSetArray();

	_, err := s.statefulSetCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("error in deletemany")
        return err
	}

	if (len(statefulSetInfo) > 0) {
		_, err2 := s.statefulSetCollection.InsertMany(context.TODO(),statefulSetInfo);
		if err2 != nil {
			fmt.Println("error in insertmany")
			return err2
		}
	}

	fmt.Println("refreshed statefulSet info")
	return nil
}

func (s *Service) initStatefulSetArray() []interface{}{
	statefulSetList, err := s.clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var statefulSets []interface{}
	for _, s := range statefulSetList.Items {
		statefulSet := object.StatefulSet {
			ObjectMeta: object.ObjectMeta{
				Name: s.Name,
				Namespace: string(s.Namespace),
				Uid: string(s.UID),
				CreationTime: s.CreationTimestamp.String(),
			},
			Replicas: int(*s.Spec.Replicas),
			ServiceName: string(s.Spec.ServiceName),
			PodManagementPolicy: string(s.Spec.PodManagementPolicy),
			CurrentReplicas: int(s.Status.CurrentReplicas),
			UpdatedReplicas: int(s.Status.UpdatedReplicas),
			ReadyReplicas: int(s.Status.ReadyReplicas),
		}
		matchLabels := make(map[string]string)
		mlb := s.Spec.Selector.MatchLabels
		for key, val := range mlb {
			matchLabels[key] = val
		}
		statefulSet.MatchLabels = matchLabels
		statefulSets = append(statefulSets, statefulSet)
	}
	return statefulSets;
}

func (s *Service) GetStatefulSetInfo(c *gin.Context) {
	cursor, err := s.statefulSetCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
		return
	}

	util.ResponseSuccess(c, results, "statefulSet")
}