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


func (s *Service) initClusterRoles() []interface{}{

	clusterRoleList, err := s.clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var clusterRoles []interface{};

	for _, r := range clusterRoleList.Items{
		clusterRole := object.ClusterRole {
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
		clusterRole.Rules = rules;

		selectors := make(map[string]string)
		for _, s := range r.AggregationRule.ClusterRoleSelectors {
			for k, v := range s.MatchLabels {
				selectors[k] = v
			}
		}
		clusterRole.ClusterRoleSelectors = selectors

		clusterRoles = append(clusterRoles, clusterRole);
	}

	return clusterRoles;
}

func (s *Service) GetClusterRoleInfo(c *gin.Context) {
	cursor, err := s.clusterRoleCollection.Find(context.TODO(), bson.D{})
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

	util.ResponseSuccess(c, results, "clusterRole")
}


func (s *Service) refreshClusterRoleInfo() error {
	clusterRoleInfo := s.initClusterRoles();

	_, err := s.clusterRoleCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err = s.clusterRoleCollection.InsertMany(context.TODO(),clusterRoleInfo);
	if err != nil {
		return err
	}
	
	fmt.Println("refreshed clusterRole info")
	return nil
}