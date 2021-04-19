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

func (s *Service) initRoleBindings() []interface{} {

	roleBindingList, err := s.clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var rbs []interface{}

	for _, r := range roleBindingList.Items {
		rb := object.RoleBinding{
			ObjectMeta: object.ObjectMeta{
				Name:         r.Name,
				Namespace:    string(r.Namespace),
				Uid:          string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
			RoleRef: object.RoleRef{
				Kind:     r.RoleRef.Kind,
				APIGroup: r.RoleRef.APIGroup,
				Name:     r.RoleRef.Name,
			},
		}
		var subjects []object.Subject
		for _, s := range r.Subjects {
			sb := object.Subject{
				Kind:      s.Kind,
				APIGroup:  s.APIGroup,
				Name:      s.Name,
				Namespace: s.Namespace,
			}
			subjects = append(subjects, sb)
		}
		rb.Subjects = subjects

		rbs = append(rbs, rb)
	}

	return rbs
}

func (s *Service) GetRoleBindingInfo(c *gin.Context) {
	var results []object.RoleBinding
	var tmp object.RoleBinding
	cursor, err := s.roleBindingCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "roleBinding")
}

func (s *Service) refreshRoleBindingInfo() error {
	roleBindingInfo := s.initRoleBindings()

	_, err := s.roleBindingCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err = s.roleBindingCollection.InsertMany(context.TODO(), roleBindingInfo)
	if err != nil {
		return err
	}

	fmt.Println("refreshed roleBinding info")
	return nil
}
