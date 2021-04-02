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
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/huantingwei/fyp/util"
)

const loginScript = "./backend/login/loginAPI.sh"
const urlFile = "./url.txt"
const tokenFile = "./token.txt"
const clientFile = "./client.txt"

type Service struct {
	cmdStdinPipe io.WriteCloser
	userCollection *mongo.Collection
}

type Project struct {
	ClusterName string 	`json:"clusterName"`
	ProjectName string 	`json:"projectName"`
	ZoneName    string 	`json:"zoneName"`
	Token		string	`json:"token"`
	CredPath	string	`json:"credPath"`
}

type GoogleCode struct {
	Code string `json:"code"`
}


func NewService(r *gin.RouterGroup, db util.Database) {
	s := &Service{
		userCollection: db.Handle.Collection("user"),
	}

	r = r.Group("/login")

	r.POST("/authenticate", s.Authenticate)
	r.POST("/verifyCode", s.VerifyCode)
	r.GET("/project", s.GetProject)
}

func (s *Service) Authenticate(c *gin.Context) {

	var err error
	var p Project
	c.ShouldBindJSON(&p)

	cmd := exec.Command(loginScript, p.ClusterName, p.ProjectName, p.ZoneName, p.CredPath)

	fmt.Println("running login script")
	// Start(): does not wait for script to complete
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error in running login script: %v\n", err)
		panic(err)
	}

	var url string
	fmt.Println("start looping for url")
	for {
		if f, err := os.Stat(urlFile); err == nil {
			if f.Size() > 0 {
				file, err := os.Open(urlFile)
				if err != nil {
					fmt.Printf(err.Error())
					util.ResponseError(c, err)
					return
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				scanner.Scan()
				url = scanner.Text()
				break
			}
		}
	}

	if err = s.newProject(p); err != nil {
		// revoke
		revoke := exec.Command("gcloud", "auth", "revoke", "--all")
		revoke.Run()
		util.ResponseError(c, err)
	}

	util.ResponseSuccess(c, url, "login")
	return
}

func (s *Service) VerifyCode(c *gin.Context) {
	var verificationCode GoogleCode
	c.ShouldBindJSON(&verificationCode)

	f, _ := os.Create(tokenFile)
	_, err := f.WriteString(verificationCode.Code + "\n")
	f.Close()
	if err != nil {
		util.ResponseError(c, err)
        return
	}

	// return project info
	p, err := s.getProject()
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, p, "project verified")
}

func (s *Service) GetProject(c *gin.Context){
	p, err := s.getProject()
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, p, "project info")
}

func (s *Service) newProject(p Project) (err error) {
	// generate token
	token, err := generateToken(p.ClusterName)
	if err != nil {
		return err
	}
	p.Token = token

	// delete any previous records
	_, err = s.userCollection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	// insert a new one
	_, err = s.userCollection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	// create local file for GetGCPClusterManagementClient()
	f, _ := os.Create(clientFile)
	_, err = f.WriteString(p.CredPath + "\n")
	if err != nil {
		return err
	}

	// projects/$PROJECTNAME/locations/$ZONENAME/clusters/$CLUSTERNAME
	clusterMeta := "projects/" + p.ProjectName + "/locations/" + p.ZoneName + "/clusters/" + p.ClusterName + "\n"
	_, err = f.WriteString(clusterMeta)
	if err != nil {
		return err
	}
	f.Close()
	
	return nil
}

func (s *Service) getProject() (Project, error) {
	var p Project
	err :=  s.userCollection.FindOne(context.Background(), bson.M{}).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (s *Service) asyncBash() {
	cmd := exec.Command(loginScript)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	s.cmdStdinPipe = stdin

	cmd.Run()
}

func generateToken(base string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(base), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}