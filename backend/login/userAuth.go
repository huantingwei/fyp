package login

import (
	"context"
	"fmt"
	"io"

	//"io/ioutil"
	"bufio"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/huantingwei/fyp/util"
)

type Service struct {
	cmdStdinPipe io.WriteCloser
	userCollection *mongo.Collection
}

type Project struct {
	ClusterName string `json:"clusterName"`
	ProjectName string `json:"projectName"`
	ZoneName    string `json:"zoneName"`
}

type GoogleCode struct {
	Code string `json:"code"`
}


func NewService(r *gin.RouterGroup, db util.Database) {
	s := &Service{
		userCollection: db.Handle.Collection("user"),
	}

	go s.asyncBash()

	r.POST("/login", s.postProjectInfo)
	r.POST("/googleAuth", s.postGoogleCode)
	r.GET("/project", s.GetProject)
	r.POST("/project", s.NewProject)

}

func (s *Service) asyncBash() {
	cmd := exec.Command("./login/login.sh")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	s.cmdStdinPipe = stdin

	cmd.Run()
}

func (s *Service) NewProject(c *gin.Context){
	var p Project
	c.ShouldBindJSON(&p)
	var err error
	// delete any previous records
	_, err = s.userCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	// insert a new one
	_, err = s.userCollection.InsertOne(context.Background(), p)
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, p, "new project")
}

func (s *Service) GetProject(c *gin.Context){
	var p *Project
	err :=  s.userCollection.FindOne(context.Background(), bson.M{}).Decode(&p)
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, p, "project info")
}

func (s *Service) postProjectInfo(c *gin.Context) {

	var p Project
	c.ShouldBindJSON(&p)
	/*
	// delete any previous records
	_, err := s.userCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		util.ResponseError(c, err)
	}

	// insert a new one
	_, err := s.userCollection.InsertOne(context.Background(), *p)
	if err != nil {
		util.ResponseError(c, err)
	}
	*/

	io.WriteString(s.cmdStdinPipe, p.ClusterName+"\n")
	io.WriteString(s.cmdStdinPipe, p.ProjectName+"\n")
	io.WriteString(s.cmdStdinPipe, p.ZoneName+"\n")

	urlChan := make(chan string)

	go func() {
		urlChan <- readURLFromFile()
	}()

	url := <-urlChan
	util.ResponseSuccess(c, url, "login")
}

func readURLFromFile() string {
	for {
		if f, err := os.Stat("./url.txt"); err == nil {
			if f.Size() > 0 {
				file, err2 := os.Open("./url.txt")
				if err2 != nil {
					fmt.Printf(err2.Error())
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				scanner.Scan()
				return scanner.Text()
			}
		}
	}
}

func (s *Service) postGoogleCode(c *gin.Context) {
	var verificationCode GoogleCode
	c.ShouldBindJSON(&verificationCode)

	f, _ := os.Create("./token.txt")
	_, err := f.WriteString(verificationCode.Code + "\n")
	f.Close()
	if err != nil {
		util.ResponseError(c, err)
        return
	}
	util.ResponseSuccess(c, verificationCode, "verification")
}
