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

func (s *Service) initNetworkPolicies() []interface{} {

	networkPolicyList, err := s.clientset.NetworkingV1().NetworkPolicies("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var nps []interface{}

	for _, n := range networkPolicyList.Items {
		np := object.NetworkPolicy{
			ObjectMeta: object.ObjectMeta{
				Name:         n.Name,
				Namespace:    string(n.Namespace),
				Uid:          string(n.UID),
				CreationTime: n.CreationTimestamp.String(),
			},
		}

		var egRules []object.NetworkPolicyEgressRule
		var inRules []object.NetworkPolicyIngressRule
		var egPorts []object.NetworkPolicyPort
		var egTo []object.NetworkPolicyPeer
		var inPorts []object.NetworkPolicyPort
		var inFrom []object.NetworkPolicyPeer
		var policyTypes = make([]string, len(n.Spec.PolicyTypes))

		// EgressRule
		for _, egr := range n.Spec.Egress {

			r := object.NetworkPolicyEgressRule{}
			for _, p := range egr.Ports {
				port := object.NetworkPolicyPort{
					Port:     int(p.Port.IntVal),
					Protocol: p.Protocol,
				}
				egPorts = append(egPorts, port)
			}
			for _, t := range egr.To {
				to := object.NetworkPolicyPeer{}

				if t.IPBlock != nil {
					to.CIDR = t.IPBlock.CIDR
					except := make([]string, len(t.IPBlock.Except))
					for _, e := range t.IPBlock.Except {
						except = append(except, e)
					}
					to.Except = except
				}

				npLabels := make(map[string]string)
				if t.NamespaceSelector != nil {
					for key, val := range t.NamespaceSelector.MatchLabels {
						npLabels[key] = val
					}
				}

				podLabels := make(map[string]string)
				if t.NamespaceSelector != nil {
					for key, val := range t.PodSelector.MatchLabels {
						podLabels[key] = val
					}
				}
				to.NamespaceSelector = npLabels
				to.PodSelector = podLabels

				egTo = append(egTo, to)
			}

			r.Ports = egPorts
			r.To = egTo
			egRules = append(egRules, r)
		}

		// IngressRule
		for _, inr := range n.Spec.Ingress {

			r := object.NetworkPolicyIngressRule{}
			for _, p := range inr.Ports {
				port := object.NetworkPolicyPort{
					Port:     int(p.Port.IntVal),
					Protocol: p.Protocol,
				}
				inPorts = append(inPorts, port)
			}
			for _, f := range inr.From {
				from := object.NetworkPolicyPeer{}

				if f.IPBlock != nil {
					from.CIDR = f.IPBlock.CIDR
					except := make([]string, len(f.IPBlock.Except))
					for _, e := range f.IPBlock.Except {
						except = append(except, e)
					}
					from.Except = except
				}

				npLabels := make(map[string]string)
				if f.NamespaceSelector != nil {
					for key, val := range f.NamespaceSelector.MatchLabels {
						npLabels[key] = val
					}
				}
				podLabels := make(map[string]string)
				if f.PodSelector != nil {
					for key, val := range f.PodSelector.MatchLabels {
						podLabels[key] = val
					}
				}
				from.NamespaceSelector = npLabels
				from.PodSelector = podLabels

				inFrom = append(inFrom, from)
			}

			r.Ports = inPorts
			r.From = inFrom
			inRules = append(inRules, r)
		}

		for _, pt := range n.Spec.PolicyTypes {
			policyTypes = append(policyTypes, string(pt))
		}

		np.NetworkPolicyEgressRule = egRules
		np.NetworkPolicyIngressRule = inRules
		np.PolicyTypes = policyTypes
		nps = append(nps, np)
	}

	return nps
}

func (s *Service) GetNetworkPolicyInfo(c *gin.Context) {
	var results []object.NetworkPolicy
	cursor, err := s.networkPolicyCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.NetworkPolicy
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "networkPolicy")
}

func (s *Service) refreshNetworkPolicyInfo() error {
	networkPolicyInfo := s.initNetworkPolicies()

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
