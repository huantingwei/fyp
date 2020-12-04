package kubescore

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"
)

type Service struct {
	kubescoreCollection *mongo.Collection
}

func NewService(r *gin.RouterGroup, db util.Database) {
	s := &Service{
		kubescoreCollection: db.Handle.Collection("kubescore"),
	}

	r = r.Group("/kubescore")

	r.GET("/result", s.GetAllKubeScore)
}
