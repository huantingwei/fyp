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

	r.GET("/get", s.GetKubescore)
	r.GET("/new", s.NewKubescore)
	r.GET("/list", s.ListKubescore)
	r.POST("/delete", s.DeleteKubescore)
	r.GET("/interactive/get", s.GetInteractiveResult)
	r.POST("/interactive/upload", s.UploadInteractiveFile)
}
