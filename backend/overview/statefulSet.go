package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) refreshStatefulSetInfo(c *gin.Context){
	statefulSetInfo := s.initStatefulSetArray();

	_, err := s.statefulSetCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	_, err2 := s.statefulSetCollection.InsertMany(context.TODO(),statefulSetInfo);
	if err2 != nil {
		util.ResponseError(c, err2)
	}

	fmt.Println("refreshed statefulSet info")
}

func (s *Service) initStatefulSetArray() []interface{}{
	statefulSetList, err := s.clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var statefulSets []interface{};
	for _, s := range statefulSetList.Items {
		statefulSet := s.DeepCopy()
		statefulSets = append(statefulSets, statefulSet)
	}
	

	return statefulSets;
}

func (s *Service) GetStatefulSetInfo(c *gin.Context) {
	cursor, err := s.statefulSetCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err2 := cursor.All(context.TODO(), &results); err2 != nil {
		util.ResponseError(c, err2)
	}

	util.ResponseSuccess(c, results, "statefulSet")
}