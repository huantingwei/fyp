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

func (s *Service) initClusterRoles() []interface{} {

	clusterRoleList, err := s.clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var clusterRoles []interface{}

	for _, r := range clusterRoleList.Items {
		clusterRole := object.ClusterRole{
			ObjectMeta: object.ObjectMeta{
				Name:         r.Name,
				Namespace:    string(r.Namespace),
				Uid:          string(r.UID),
				CreationTime: r.CreationTimestamp.String(),
			},
		}

		var rules []object.PolicyRule
		for _, rule := range r.Rules {
			var apiGroups []string
			var nonResourceUrls []string
			var resourceNames []string
			var resources []string
			var verbs []string
			apiGroups = append(apiGroups, rule.APIGroups...)
			nonResourceUrls = append(nonResourceUrls, rule.NonResourceURLs...)
			resourceNames = append(resourceNames, rule.ResourceNames...)
			resources = append(apiGroups, rule.Resources...)
			verbs = append(apiGroups, rule.Verbs...)

			pr := object.PolicyRule{
				APIGroups:       apiGroups,
				NonResourceURLs: nonResourceUrls,
				ResourceNames:   resourceNames,
				Resources:       resources,
				Verbs:           verbs,
			}
			rules = append(rules, pr)
		}
		clusterRole.Rules = rules

		selectors := make(map[string]string)
		if r.AggregationRule != nil {
			for _, s := range r.AggregationRule.ClusterRoleSelectors {
				for k, v := range s.MatchLabels {
					selectors[k] = v
				}
			}
		}
		clusterRole.ClusterRoleSelectors = selectors

		clusterRoles = append(clusterRoles, clusterRole)
	}

	return clusterRoles
}

func (s *Service) GetClusterRoleInfo(c *gin.Context) {
	var results []object.ClusterRole
	var tmp object.ClusterRole
	cursor, err := s.clusterRoleCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "clusterRole")
}

func (s *Service) refreshClusterRoleInfo() error {
	clusterRoleInfo := s.initClusterRoles()

	_, err := s.clusterRoleCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	if len(clusterRoleInfo) > 0 {
		_, err = s.clusterRoleCollection.InsertMany(context.TODO(), clusterRoleInfo)
		if err != nil {
			return err
		}

		fmt.Println("refreshed clusterRole info")
	}
	return nil
}
