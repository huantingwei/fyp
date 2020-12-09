package kubebench

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/huantingwei/fyp/util"
)

type Service struct {
	kubebenchCollection *mongo.Collection
}

func NewService(r *gin.RouterGroup, db util.Database) {
	s := &Service{
		kubebenchCollection: db.Handle.Collection("kubebench"),
	}
	r = r.Group("/kubebench")

	r.GET("/get", s.GetKubebench)
	r.GET("/new", s.NewKubebench)

}
