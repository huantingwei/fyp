package kubescore

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/huantingwei/fyp/util"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	dbName = "fyp"
	coll   = "kubescore"
	// windows
	// resultFile = "kubescore\\result.json"
	// scriptFile = "kubescore\\kubescore.sh"

	// linux
	resultFile = "./backend/kubescore/res.json"
	runScript = "./backend/kubescore/run.sh"
)

type Check struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	TargetType string `json:"target_type"`
	Comment    string `json:"comment"`
	Optional   bool   `json:"optional"`
}

type ScoredObject struct {
	ObjectName string            `json:"object_name"`
	TypeMeta   metav1.TypeMeta   `json:"type_meta"`
	ObjectMeta metav1.ObjectMeta `json:"object_meta"`
	Checks     []TestScore       `json:"checks"`
	FileName   string            `json:"file_name"`
	FileRow    int               `json:"file_row"`
}

type TestScore struct {
	Check Check `json:"check"`
	// Grade    scorecard.Grade    `json:"grade"`
	Skipped  bool               `json:"skipped"`
	Comments []TestScoreComment `json:"comments"`
}

type TestScoreComment struct {
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type KubeScore struct {
	ID         primitive.ObjectID `json:"id"`
	CreateTime string             `json:"createTime"`
	ScoredObjects []ScoredObject `json:"kubescore"`
}


func readFile() (*KubeScore, error) {
	jsonFile, err := os.Open(resultFile)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var kbscores []ScoredObject

	err = json.Unmarshal(byteValue, &kbscores)
	if err != nil {
		fmt.Printf("Error in reading kubescore result json file: %v\n", err)
		return nil, err
	}
	id := primitive.NewObjectID()
	t := time.Now()
	fmtS := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	return &KubeScore{id, fmtS, kbscores}, nil
}

func create(kubescore *KubeScore, s *Service) (insertedID interface{}, err error) {

	_, err = s.kubescoreCollection.InsertOne(context.Background(), *kubescore)
	if err != nil {
		fmt.Printf("Error in creating kubescore result: %v\n", err)
		return "-1", err
	}

	fmt.Printf("Successful create; Inserted ID: %s\n", kubescore.ID)
	return kubescore.ID, nil

}

func read(id string, s *Service) (kubescore *KubeScore, err error) {
	// convert string to primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(id)

	err = s.kubescoreCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: oid}}).Decode(&kubescore)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successful read id: %s\n", id)
	return
}

// GetKubescore retrieve a single Kubescore object with "_id" = "id"
// Response: result of the given id
func (s *Service) GetKubescore(c *gin.Context) {
	// query parameters
	id := c.Query("id")

	var kbscore *KubeScore
	kbscore, err := read(id, s)
	if err != nil {
		fmt.Printf("Error in GetKubescore, id=%s: %v\n", id, err)
		util.ResponseError(c, err)
		return
	} else {
		util.ResponseSuccess(c, kbscore, "kubescore")
		return
	}
}

// NewKubescore run the kubescore script, read from result file, and create result object in DB
// Response: id of the result object
func (s *Service) NewKubescore(c *gin.Context) {

	var kbscore *KubeScore
	var id interface{}
	var err error

	cmd := exec.Command(runScript)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error in waiting for kubescore script executing: %v\n", err)
		goto responseError
	}

	// read from result file
	kbscore, err = readFile()
	if err != nil {
		fmt.Printf("Error in read kubescore result file: %v\n", err)
		goto responseError
	}
	// create result in DB
	id, err = create(kbscore, s)
	if err != nil {
		fmt.Printf("Error in create kubescore result in DB: %v\n", err)
		goto responseError
	}
	util.ResponseSuccess(c, id, "kubescore")
	return

responseError:
	util.ResponseError(c, err)
	return

}
// ListKubescore list all results
func (s *Service) ListKubescore(c *gin.Context) {
	var kcs []KubeScore
	cursor, err := s.kubescoreCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("Error in list Kubescore results: %v\n", err)
		util.ResponseError(c, err)
		return
	}
	if err = cursor.All(context.TODO(), &kcs); err != nil {
		fmt.Printf("Error in decoding Kubescore result list: %v\n", err)
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, kcs, "kubescore")
}

func (s *Service) DeleteKubescore(c *gin.Context) {
	var tmp KubeScore
	c.ShouldBindJSON(&tmp)

	res, err := s.kubescoreCollection.DeleteOne(context.TODO(), bson.M{"id": tmp.ID})
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, int(res.DeletedCount), "kubescore")
}
