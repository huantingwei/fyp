package network

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"

	kube "k8s.io/client-go/kubernetes"
)

type Service struct {
	networkCollection    	*mongo.Collection
	networkGraphCollection	*mongo.Collection
	clientset            	*kube.Clientset
}

func NewService(r *gin.RouterGroup, db util.Database, client *kube.Clientset) {
	s := &Service{
		networkCollection:    		db.Handle.Collection("network"),
		networkGraphCollection:		db.Handle.Collection("networkGraph"),
		clientset:            		client,
	}
	
	// initialize graph
	//s.insertGraph()

	r = r.Group("/network")

	r.GET("/graph", s.GetGraph)
	r.GET("/namespace", s.GetNamespace)
	r.POST("/graph", s.RefreshGraph)
}
