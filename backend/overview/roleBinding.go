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


func (s *Service) initRoleBindings() []interface{}{

	roleBindingList, err := s.clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var rbs []interface{};

	for _, r := range roleBindingList.Items{
		rb := object.RoleBinding {
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

func (s *Service) GetRoleBindingInfo(c *gin.Context) {
	cursor, err := s.roleBindingCollection.Find(context.TODO(), bson.D{})
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

	util.ResponseSuccess(c, results, "roleBinding")
}


func (s *Service) refreshRoleBindingInfo(c *gin.Context){
	roleBindingInfo := s.initRoleBindings();

	_, err := s.roleBindingCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	_, err = s.roleBindingCollection.InsertMany(context.TODO(),roleBindingInfo);
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	
	fmt.Println("refreshed roleBinding info")
}