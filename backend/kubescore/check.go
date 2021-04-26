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
	// linux
	resultFile            = "./backend/kubescore/res.json"
	runScript             = "./backend/kubescore/run.sh"
	interactiveScript     = "./backend/kubescore/interactive.sh"
	interactiveFileInput  = "./backend/kubescore/interactive-in.yaml"
	interactiveFileOutput = "./backend/kubescore/interactive-out.json"
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
	ID            primitive.ObjectID `json:"id"`
	CreateTime    string             `json:"createTime"`
	ScoredObjects []ScoredObject     `json:"kubescore"`
}

type FileInput struct {
	Content string `json:"fileContent"`
}

func readFile(file string) (*KubeScore, error) {
	jsonFile, err := os.Open(file)
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
	var err error

	// delete previous records
	err = os.Remove(resultFile)
	// ignore error if file does not exist

	_, err = s.kubescoreCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	cmd := exec.Command(runScript)
	// Start(): no wait
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error in executing for kubescore script executing: %v\n", err)
		util.ResponseError(c, err)
		return
	}
	// successful execution
	util.ResponseSuccess(c, "start scanning...", "kubescore")
	return
}

// ListKubescore list all results
func (s *Service) ListKubescore(c *gin.Context) {

	var kcs []KubeScore
	var kbscore *KubeScore
	var err error

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

	// no record in db
	if kcs == nil || len(kcs) == 0 {
		// read from result file
		kbscore, err = readFile(resultFile)
		if err != nil {
			fmt.Printf("Error in read kubescore result file: %v\n", err)
			goto empty
		}
		// create result in DB
		_, err := create(kbscore, s)
		if err != nil {
			fmt.Printf("Error in create kubescore result in DB: %v\n", err)
			goto empty
		}
		kcs = append(kcs, *kbscore)
	}

	util.ResponseSuccess(c, kcs, "kubescore")
	return

empty:
	var empty []interface{}
	util.ResponseSuccess(c, empty, "kubescore")
	return
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

func (s *Service) UploadInteractiveFile(c *gin.Context) {
	var err error
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("cannot read formfile")
		util.ResponseError(c, err)
		return
	}
	err = c.SaveUploadedFile(file, interactiveFileInput)
	if err != nil {
		fmt.Println("cannot upload file")
		util.ResponseError(c, err)
		return
	}

	cmd := exec.Command(interactiveScript)
	err = cmd.Start()
	if err != nil {
		fmt.Println("cannot execute kubescore")
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, "successful upload", "interactive kubescore")
}

func (s *Service) GetInteractiveResult(c *gin.Context) {
	var err error
	var kbscore *KubeScore
	kbscore, err = readFile(interactiveFileOutput)
	if err != nil {
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, *kbscore, "interactive kubescore")
}
