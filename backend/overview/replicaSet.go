package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshReplicaSetInfo(c *gin.Context){
	replicaSetInfo := s.initReplicaSetArray();

	_, err := s.replicaSetCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	_, err2 := s.replicaSetCollection.InsertMany(context.TODO(),replicaSetInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
	}

	fmt.Println("refreshed replicaSet info")
}

func (s *Service) initReplicaSetArray() []interface{}{
	replicaSetList, err := s.clientset.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var replicaSets []interface{};
	for _, s := range replicaSetList.Items {
		replicaSet := s.DeepCopy()
		replicaSets = append(replicaSets, replicaSet)
	}
	

	return replicaSets;
}

func (s *Service) GetReplicaSetInfo(c *gin.Context) {
	cursor, err := s.replicaSetCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
	}

	util.ResponseSuccess(c, results, "replicaSet")
}