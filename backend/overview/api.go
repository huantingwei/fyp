package overview

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"

	kube "k8s.io/client-go/kubernetes"
)

type Service struct {
	clusterCollection    	*mongo.Collection
	deploymentCollection 	*mongo.Collection
	nodeCollection       	*mongo.Collection
	nodepoolCollection   	*mongo.Collection
	podCollection        	*mongo.Collection
	serviceCollection    	*mongo.Collection
	roleCollection		 	*mongo.Collection
	roleBindingCollection	*mongo.Collection
	networkPolicyCollection *mongo.Collection
	statefulSetCollection	*mongo.Collection
	replicaSetCollection	*mongo.Collection
	clientset            	*kube.Clientset
}

func NewService(r *gin.RouterGroup, db util.Database, client *kube.Clientset) {
	s := &Service{
		clusterCollection:    		db.Handle.Collection("cluster"),
		deploymentCollection: 		db.Handle.Collection("deployment"),
		nodeCollection:       		db.Handle.Collection("node"),
		nodepoolCollection:   		db.Handle.Collection("nodepool"),
		podCollection:        		db.Handle.Collection("pod"),
		serviceCollection:    		db.Handle.Collection("service"),
		roleCollection:		  		db.Handle.Collection("role"),
		roleBindingCollection:		db.Handle.Collection("roleBinding"),
		networkPolicyCollection:	db.Handle.Collection("networkPolicy"),
		statefulSetCollection: 		db.Handle.Collection("statefulSet"),
		replicaSetCollection:		db.Handle.Collection("replicaSet"),
		clientset:            		client,
	}

	// initialize cluser data
	s.init()

	r = r.Group("/overview")

	r.GET("/cluster", s.GetClusterInfo)
	r.GET("/nodepool", s.GetNodepoolInfo)
	r.GET("/deployment", s.GetDeploymentInfo)
	r.GET("/node", s.GetNodeInfo)
	r.GET("/pod", s.GetPodInfo)
	r.GET("/service", s.GetServiceInfo)
	r.GET("/role", s.GetRoleInfo)
	r.GET("/roleBinding", s.GetRoleBindingInfo)
	r.GET("/networkPolicy", s.GetNetworkPolicyInfo)
	r.GET("/statefulSet", s.GetStatefulSetInfo)
	r.GET("/replicaSet", s.GetReplicaSetInfo)
	
	r.POST("/new", s.Refresh)
}

func (s *Service) init(){
	// TODO: remove c gin.Context
//	s.refreshClusterInfo(c)
//	s.refreshNodepoolInfo(c)
//	s.refreshDeploymentInfo(c)
//	s.refreshPodInfo(c)
//	s.refreshServiceInfo(c)
//	s.refreshNodeInfo(c)
//	s.refreshRoleInfo(c)
//	s.refreshRoleBindingInfo(c)
//	s.refreshNetworkPolicyInfo(c)
//	s.refreshStatefulSetInfo(c)
//	s.refreshReplicaSetInfo(c)
	return
}

func (s *Service) Refresh(c *gin.Context) {
	s.refreshClusterInfo(c)
	s.refreshNodepoolInfo(c)
	s.refreshDeploymentInfo(c)
	s.refreshPodInfo(c)
	s.refreshServiceInfo(c)
	s.refreshNodeInfo(c)
	s.refreshRoleInfo(c)
	s.refreshRoleBindingInfo(c)
	s.refreshNetworkPolicyInfo(c)
	s.refreshStatefulSetInfo(c)
	s.refreshReplicaSetInfo(c)

	var data interface{}
	util.ResponseSuccess(c, data, "refreshed all kube resources")
}
