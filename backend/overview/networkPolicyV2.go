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

func (s *Service) initNetworkPoliciesV2() []interface{} {

	networkPolicyList, err := s.clientset.NetworkingV1().NetworkPolicies("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var nps []interface{}

	for _, n := range networkPolicyList.Items {
		np := object.NetworkPolicyV2{
			ObjectMeta: object.ObjectMeta{
				Name:         n.Name,
				Namespace:    string(n.Namespace),
				Uid:          string(n.UID),
				CreationTime: n.CreationTimestamp.String(),
			},
		}

		var egRules []object.NetworkPolicyEgressRuleV2
		var inRules []object.NetworkPolicyIngressRuleV2
		egPorts := make(map[string]interface{})
		egTo := make(map[string]string)
		inPorts := make(map[string]interface{})
		inFrom := make(map[string]string)
		policyTypes := ""

		// EgressRule
		for _, egr := range n.Spec.Egress {

			r := object.NetworkPolicyEgressRuleV2{}

			for _, p := range egr.Ports {
				// big bug here: same port has different protocol or vice versa
				egPorts[p.Port.String()] = p.Protocol
			}
			for _, t := range egr.To {
				if t.IPBlock != nil {
					egTo["CIDR"] = t.IPBlock.CIDR
					egTo["Except"] = ""
					for _, e := range t.IPBlock.Except {
						egTo["Except"] += e
					}
				}

				egTo["NamespaceSelector"] = ""
				if t.NamespaceSelector != nil {
					for key, val := range t.NamespaceSelector.MatchLabels {
						egTo["NamespaceSelector"] += (key + ":" + val + ", ")
					}
				}

				egTo["PodSelector"] = ""
				if t.PodSelector != nil {
					for key, val := range t.PodSelector.MatchLabels {
						egTo["PodSelector"] += (key + ":" + val + ", ")
					}
				}
			}

			r.Ports = egPorts
			r.To = egTo
			egRules = append(egRules, r)
		}

		// IngressRule
		for _, ingr := range n.Spec.Ingress {

			r := object.NetworkPolicyIngressRuleV2{}
			for _, p := range ingr.Ports {
				// big bug here: same port has different protocol or vice versa
				inPorts[p.Port.String()] = p.Protocol
			}
			for _, t := range ingr.From {
				if t.IPBlock != nil {
					inFrom["CIDR"] = t.IPBlock.CIDR
					inFrom["Except"] = ""
					for _, e := range t.IPBlock.Except {
						inFrom["Except"] += e
					}
				}

				inFrom["NamespaceSelector"] = ""
				if t.NamespaceSelector != nil {
					for key, val := range t.NamespaceSelector.MatchLabels {
						inFrom["NamespaceSelector"] += (key + ":" + val + ", ")
					}
				}

				inFrom["PodSelector"] = ""
				if t.PodSelector != nil {
					for key, val := range t.PodSelector.MatchLabels {
						inFrom["PodSelector"] += (key + ":" + val + ", ")
					}
				}
			}

			r.Ports = inPorts
			r.From = inFrom
			inRules = append(inRules, r)
		}

		for i, pt := range n.Spec.PolicyTypes {
			if i != 0 {
				policyTypes += ", "
			}
			policyTypes += string(pt)
		}

		np.NetworkPolicyEgressRule = egRules
		np.NetworkPolicyIngressRule = inRules
		np.PolicyTypes = policyTypes
		nps = append(nps, np)
	}

	return nps
}

func (s *Service) GetNetworkPolicyInfoV2(c *gin.Context) {
	var results []object.NetworkPolicyV2
	cursor, err := s.networkPolicyCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.NetworkPolicyV2
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "networkPolicy")
}

func (s *Service) refreshNetworkPolicyInfoV2() error {
	networkPolicyInfo := s.initNetworkPoliciesV2()

	_, err := s.networkPolicyCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	if len(networkPolicyInfo) > 0 {
		_, err = s.networkPolicyCollection.InsertMany(context.TODO(), networkPolicyInfo)
		if err != nil {
			return err
		}
	}

	fmt.Println("refreshed networkPolicy info")
	return nil
}
