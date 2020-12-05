package kubebench

import (
	"log"

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
	err := readWriteResult()
	if err != nil {
		log.Fatal(err)
	}
	r = r.Group("/kubebench")
	r.GET("/result/get", s.GetKubebenchResult)

}
