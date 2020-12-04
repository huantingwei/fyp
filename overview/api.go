package overview

import(
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"
)

type Service struct {
	clusterCollection     *mongo.Collection
	deploymentCollection  *mongo.Collection
	nodeCollection        *mongo.Collection
	nodepoolCollection    *mongo.Collection
	podCollection         *mongo.Collection
	serviceCollection     *mongo.Collection
}

func NewService(r *gin.RouterGroup, db util.Database){
	s := &Service{
		clusterCollection: db.Handle.Collection("cluster"),
		deploymentCollection: db.Handle.Collection("deployment"),
		nodeCollection: db.Handle.Collection("node"),
		nodepoolCollection: db.Handle.Collection("nodepool"),
		podCollection: db.Handle.Collection("pod"),
		serviceCollection: db.Handle.Collection("service"),
	}

	r = r.Group("/overview");

	r.GET("/cluster", s.GetClusterInfo);
	r.GET("/nodepool", s.GetNodepoolInfo);
	r.GET("/deployment", s.GetDeploymentInfo);
	r.GET("/node", s.GetNodeInfo);
	r.GET("/pod", s.GetPodInfo);
	r.GET("/service", s.GetServiceInfo);
}