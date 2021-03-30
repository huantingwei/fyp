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


func (s *Service) initRoles() []interface{}{

	roleList, err := s.clientset.RbacV1().Roles("").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var roles []interface{};

	for _, r := range roleList.Items{
		role := object.Role {
			ObjectMeta: object.ObjectMeta{
				Name: r.Name,
				Namespace: string(r.Namespace),
				Uid: string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
		}

		var rules []object.PolicyRule;
		for _, rule := range r.Rules {
			var apiGroups []string;
			var nonResourceUrls []string;
			var resourceNames []string;
			var resources []string;
			var verbs []string;
			for _, i := range rule.APIGroups {
				apiGroups = append(apiGroups, i);
			}
			for _, i := range rule.NonResourceURLs {
				nonResourceUrls = append(nonResourceUrls, i);
			}
			for _, i := range rule.ResourceNames {
				resourceNames = append(resourceNames, i);
			}
			for _, i := range rule.Resources {
				resources = append(resources, i);
			}
			for _, i := range rule.Verbs {
				verbs = append(verbs, i);
			}

			pr := object.PolicyRule {
				APIGroups: apiGroups,
				NonResourceURLs: nonResourceUrls,
				ResourceNames: resourceNames,
				Resources: resources,
				Verbs: verbs,
			}
			rules = append(rules, pr);
		}
		role.Rules = rules;
		roles = append(roles, role);
	}

	return roles;
}

func (s *Service) GetRoleInfo(c *gin.Context) {
	cursor, err := s.roleCollection.Find(context.TODO(), bson.D{})
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

	util.ResponseSuccess(c, results, "role")
}


func (s *Service) refreshRoleInfo() error {
	roleInfo := s.initRoles();

	_, err := s.roleCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err = s.roleCollection.InsertMany(context.TODO(),roleInfo);
	if err != nil {
		return err
	}
	
	fmt.Println("refreshed role info")
	return nil
}