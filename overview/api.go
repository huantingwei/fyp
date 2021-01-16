package overview

import(
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"

	kube "k8s.io/client-go/kubernetes"
)

type Service struct {
	clusterCollection     *mongo.Collection
	deploymentCollection  *mongo.Collection
	nodeCollection        *mongo.Collection
	nodepoolCollection    *mongo.Collection
	podCollection         *mongo.Collection
	serviceCollection     *mongo.Collection
	clientset			  *kube.Clientset
}

func NewService(r *gin.RouterGroup, db util.Database){
	s := &Service{
		clusterCollection: db.Handle.Collection("cluster"),
		deploymentCollection: db.Handle.Collection("deployment"),
		nodeCollection: db.Handle.Collection("node"),
		nodepoolCollection: db.Handle.Collection("nodepool"),
		podCollection: db.Handle.Collection("pod"),
		serviceCollection: db.Handle.Collection("service"),
		clientset: util.GetKubeClientSet(),
	}

	r = r.Group("/overview");

	r.GET("/cluster", s.GetClusterInfo);
	r.GET("/nodepool", s.GetNodepoolInfo);
	r.GET("/deployment", s.GetDeploymentInfo);
	r.GET("/node", s.GetNodeInfo);
	r.GET("/pod", s.GetPodInfo);
	r.GET("/service", s.GetServiceInfo);

	r.POST("/new", s.Refresh);
}

func (s *Service) Refresh(c *gin.Context){
	s.refreshClusterInfo(c);
	s.refreshNodepoolInfo(c);
	s.refreshDeploymentInfo(c);
	s.refreshPodInfo(c);
	s.refreshServiceInfo(c);
	s.refreshNodeInfo(c);

	var data interface{};
	util.ResponseSuccess(c, data, "refreshed all kube resources");
}