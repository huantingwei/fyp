package overview

import (
	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"github.com/huantingwei/fyp/object"
	"github.com/huantingwei/fyp/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func (s *Service) initNetworkPolicies() []interface{}{

	networkPolicyList, err := s.clientset.NetworkingV1().NetworkPolicies("default").List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		panic(err.Error())
	}

	var nps []interface{};

	for _, n := range networkPolicyList.Items{
		np := n.DeepCopy()
		nps = append(nps, np)
		/*
		np := object.NetworkPolicy {
			ObjectMeta: object.ObjectMeta{
				Name: n.Name,
				Namespace: string(n.Namespace),
				Uid: string(n.UID),
				CreationTime: n.CreationTimestamp.String(),
			},
		}
		fmt.Println(np)

		//var spec object.NetworkPolicySpec;
		var egRules []object.NetworkPolicyEgressRule;
		var inRules []object.NetworkPolicyIngressRule;
		var egPorts []object.NetworkPolicyPort;
		var egTo 	[]object.NetworkPolicyPeer;
		var inPorts []object.NetworkPolicyPort;
		var inFrom 	[]object.NetworkPolicyPeer;
		var npSelector object.LabelSelector 
		var podSelector object.LabelSelector 

		// var policyTypes []object.PolicyType;

		for _, egr := range n.Spec.Egress {
			fmt.Println("egr:", egr);
			fmt.Println("egr port:", egr.Ports);

			for _, p := range egr.Ports {
				port := object.NetworkPolicyPort {
					Port: p.Port,
					//Protocol: p.Protocol,
				}
				egPorts = append(egPorts, port);
			}
			for _, t := range egr.To {
				
				var matchExpressions []object.LabelSelectorRequirement
				for _, me := range t.NamespaceSelector.MatchExpressions {
					mex := object.LabelSelectorRequirement {
						Key: me.Key,
						//Operator: me.Operator,
					}
					var values []string;
					for _, vs := range me.Values {
						values = append(values, vs)
					}
					mex.Values = values
					matchExpressions = append(matchExpressions, mex)
				}
				matchLabels := make(map[string]string);
				for key, val := range t.NamespaceSelector.MatchLabels {
					matchLabels[key] = val;
				}
				npSelector.MatchExpressions = matchExpressions
				npSelector.MatchLabels = matchLabels

				matchExpressions = nil
				for _, me := range t.PodSelector.MatchExpressions {
					mex := object.LabelSelectorRequirement {
						Key: me.Key,
						//Operator: me.Operator,
					}
					var values []string;
					for _, vs := range me.Values {
						values = append(values, vs)
					}
					mex.Values = values
					matchExpressions = append(matchExpressions, mex)
				}
				matchLabels = nil
				for key, val := range t.PodSelector.MatchLabels {
					matchLabels[key] = val;
				}
				podSelector.MatchExpressions = matchExpressions
				podSelector.MatchLabels = matchLabels

				to := object.NetworkPolicyPeer {
					CIDR: t.IPBlock.CIDR,
					Except: t.IPBlock.Except,
					NamespaceSelector: npSelector,
					PodSelector: podSelector,
				}
				egTo = append(egTo, to);
			}
			egRules = append(egRules, egr);
			
			
		}
		
		for _, inr := range n.Spec.Ingress {
			fmt.Println("inr:", inr);
			fmt.Println("inr ports:", inr.Ports)
			
			for _, p := range inr.Ports {
				port := object.NetworkPolicyPort {
					Port: p.Port,
					//Protocol: p.Protocol,
				}
				inPorts = append(inPorts, port);
			}
			for _, f := range inr.From {
				
				var npSelector object.LabelSelector 
				var matchExpressions []object.LabelSelectorRequirement
				for _, me := range f.NamespaceSelector.MatchExpressions {
					mex := object.LabelSelectorRequirement {
						Key: me.Key,
						//Operator: me.Operator,
					}
					var values []string;
					for _, vs := range me.Values {
						values = append(values, vs)
					}
					mex.Values = values
					matchExpressions = append(matchExpressions, mex)
				}
				matchLabels := make(map[string]string);
				for key, val := range f.NamespaceSelector.MatchLabels {
					matchLabels[key] = val;
				}
				npSelector.MatchExpressions = matchExpressions
				npSelector.MatchLabels = matchLabels

				var podSelector object.LabelSelector 
				matchExpressions = nil
				for _, me := range f.PodSelector.MatchExpressions {
					mex := object.LabelSelectorRequirement {
						Key: me.Key,
						//Operator: me.Operator,
					}
					var values []string;
					for _, vs := range me.Values {
						values = append(values, vs)
					}
					mex.Values = values
					matchExpressions = append(matchExpressions, mex)
				}
				matchLabels = nil
				for key, val := range f.PodSelector.MatchLabels {
					matchLabels[key] = val;
				}
				podSelector.MatchExpressions = matchExpressions
				podSelector.MatchLabels = matchLabels

				
				from := object.NetworkPolicyPeer {
					CIDR: f.IPBlock.CIDR,
					Except: f.IPBlock.Except,
					NamespaceSelector: npSelector,
					PodSelector: podSelector,
				}
				inFrom = append(inFrom, from);
			}
			inRules = append(inRules, inr);
			
			
		}
		for key, val := range n.Spec.PodSelector.MatchLabels {
			podSelectorMap[key] = val;
		}

		for _, pt := range n.Spec.PolicyTypes {
			policyTypes = append(policyTypes, pt);
		}

		sp := object.NetworkPolicySpec {
			NetworkPolicyEgressRule: egRules,
			NetworkPolicyIngressRule: inRules,
			PodSelector: podSelectorMap,
			// PolicyTypes: policyTypes,
		}
		np.Spec = sp;

		nps = append(nps, np);
		*/
	}
	
	
	return nps;
}

func (s *Service) GetNetworkPolicyInfo(c *gin.Context) {
	cursor, err := s.networkPolicyCollection.Find(context.TODO(), bson.D{})
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

	util.ResponseSuccess(c, results, "networkPolicy")
}


func (s *Service) refreshNetworkPolicyInfo(c *gin.Context){
	networkPolicyInfo := s.initNetworkPolicies();

	_, err := s.networkPolicyCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	_, err = s.networkPolicyCollection.InsertMany(context.TODO(),networkPolicyInfo);
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	
	fmt.Println("refreshed networkPolicy info")
}
