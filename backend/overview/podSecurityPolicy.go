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

func (s *Service) initPodSecurityPolicies() []interface{} {

	podSecurityPolicyList, err := s.clientset.PolicyV1beta1().PodSecurityPolicies().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var psps []interface{}

	for _, p := range podSecurityPolicyList.Items {
		psp := object.PodSecurityPolicy{
			ObjectMeta: object.ObjectMeta{
				Name:         p.Name,
				Namespace:    string(p.Namespace),
				Uid:          string(p.UID),
				CreationTime: p.CreationTimestamp.String(),
			},
			HostIPC:                p.Spec.HostIPC,
			HostNetwork:            p.Spec.HostNetwork,
			HostPID:                p.Spec.HostPID,
			Privileged:             p.Spec.Privileged,
			ReadOnlyRootFileSystem: p.Spec.ReadOnlyRootFilesystem,
		}

		if p.Spec.AllowPrivilegeEscalation != nil {
			psp.AllowPrivilegeEscalation = *(p.Spec.AllowPrivilegeEscalation)
		}
		if p.Spec.DefaultAllowPrivilegeEscalation != nil {
			psp.DefaultAllowPrivilegeEscalation = *(p.Spec.DefaultAllowPrivilegeEscalation)
		}
		var allowedUnsafeSysctls = make([]string, len(p.Spec.AllowedUnsafeSysctls))
		var forbiddenSysctls = make([]string, len(p.Spec.ForbiddenSysctls))
		copy(allowedUnsafeSysctls, p.Spec.AllowedUnsafeSysctls)
		psp.AllowedUnsafeSysctls = allowedUnsafeSysctls
		copy(forbiddenSysctls, p.Spec.ForbiddenSysctls)
		psp.ForbiddenSysctls = forbiddenSysctls

		//var allowedCapabilities = make([]interface{}, len(p.Spec.AllowedCapabilities))
		//var defaultAddCapabilities = make([]interface{}, len(p.Spec.DefaultAddCapabilities))
		//var requiredDropCapabilities = make([]interface{}, len(p.Spec.RequiredDropCapabilities))
		//
		//psp.AllowedCapabilities = copy(allowedCapabilities, p.Spec.AllowedCapabilities)
		//psp.DefaultAddCapabilities = copy(defaultAddCapabilities, p.Spec.DefaultAddCapabilities)
		//psp.RequiredDropCapabilities = copy(requiredDropCapabilities, p.Spec.RequiredDropCapabilities)

		if p.Spec.AllowedHostPaths != nil {
			var allowedHostPaths = make([]string, len(p.Spec.AllowedHostPaths))
			for _, path := range p.Spec.AllowedHostPaths {
				//if(path.PathPrefix != nil){
				allowedHostPaths = append(allowedHostPaths, path.PathPrefix)
				//}
			}
		}
		if p.Spec.HostPorts != nil {
			var hostPorts []string
			for _, hostPortRange := range p.Spec.HostPorts {
				hostPorts = append(hostPorts, string(hostPortRange.Min)+":"+string(hostPortRange.Max))
			}
			psp.HostPorts = hostPorts
		}

		if p.Spec.RunAsGroup != nil {
			psp.RunAsGroup = p.Spec.RunAsGroup.Rule
		}
		//if(p.Spec.RunAsUser != nil){
		psp.RunAsUser = p.Spec.RunAsUser.Rule
		//}
		psps = append(psps, psp)
	}

	return psps
}

func (s *Service) GetPodSecurityPolicyInfo(c *gin.Context) {
	var results []object.PodSecurityPolicy
	cursor, err := s.podSecurityPolicyCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	for cursor.Next(context.TODO()) {
		var tmp object.PodSecurityPolicy
		cursor.Decode(&tmp)
		results = append(results, tmp)
	}

	util.ResponseSuccess(c, results, "podSecurityPolicy")
}

func (s *Service) refreshPodSecurityPolicyInfo() error {
	podSecurityPolicyInfo := s.initPodSecurityPolicies()

	_, err := s.podSecurityPolicyCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	_, err = s.podSecurityPolicyCollection.InsertMany(context.TODO(), podSecurityPolicyInfo)
	if err != nil {
		return err
	}

	fmt.Println("refreshed podSecurityPolicy info")
	return nil
}
