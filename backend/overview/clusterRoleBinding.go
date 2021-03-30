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


func (s *Service) initClusterRoleBindings() []interface{}{

	clusterRoleBindingList, err := s.clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var rbs []interface{};

	for _, r := range clusterRoleBindingList.Items{
		rb := object.ClusterRoleBinding {
			ObjectMeta: object.ObjectMeta{
				Name: r.Name,
				Namespace: string(r.Namespace),
				Uid: string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
			RoleRef: object.RoleRef {
				Kind: r.RoleRef.Kind,
				APIGroup: r.RoleRef.APIGroup,
				Name: r.RoleRef.Name,
			},
		}
		var subjects []object.Subject;
		for _, s := range r.Subjects {
			sb := object.Subject {
				Kind: s.Kind,
				APIGroup: s.APIGroup,
				Name: s.Name,
				Namespace: s.Namespace,
			}
			subjects = append(subjects, sb)
		}
		rb.Subjects = subjects;
		
		rbs = append(rbs, rb);
	}

	return rbs;
}

func (s *Service) GetClusterRoleBindingInfo(c *gin.Context) {
	cursor, err := s.clusterRoleBindingCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
        return
	}

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		util.ResponseError(c, err)
	}

	util.ResponseSuccess(c, results, "clusterRoleBinding")
}


func (s *Service) refreshClusterRoleBindingInfo() error {
	clusterRoleBindingInfo := s.initClusterRoleBindings();

	_, err := s.clusterRoleBindingCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}

	_, err = s.clusterRoleBindingCollection.InsertMany(context.TODO(),clusterRoleBindingInfo);
	if err != nil {
		return err
	}
	
	fmt.Println("refreshed clusterRoleBinding info")
	return nil
}