package kubebench

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
)

const (
	dbName = "fyp"
	coll   = "kubebench"
	// windows
	// resultFile = "kubebench\\result.json"
	// runScript = "kubebench\\kubebench.sh"

	// linux
	resultFile = "./kubebench/kb_output.json"
	runScript = "./kubebench/run.sh"
)

type Kubebench struct {
	ID         primitive.ObjectID `json:"id"`
	CreateTime string             `json:"createTime"`
	Controls   []Control		  `json:"Controls"`
	Totals	   Totals			  `json:"Totals"`
}

type Control struct {
	ID       string    `json:"id"`
	Version  string    `json:"version"`
	Text     string    `json:"text"`
	NodeType string    `json:"node_type"`
	Sections []Section `json:"tests"`

}
type Totals struct {

	TotalPass int `json:"total_pass"`
	TotalFail int `json:"total_fail"`
	TotalWarn int `json:"total_warn"`
	TotalInfo int `json:"total_info"`
}

type Section struct {
	Section string   `json:"section"`
	Type string `json:"type"`
	Pass    int      `json:"pass"`
	Warn    int      `json:"warn"`
	Info    int      `json:"info"`
	Desc    string   `json:"desc"`
	Results []Result `json:"results"`
}

type Result struct {
	TestNumber     string   `json:"test_number"`
	TestDesc       string   `json:"test_desc"`
	Audit          string   `json:"audit"`
	AuditEnv       string   `json:"AuditEnv"`
	AuditConfig    string   `json:"AuditConfig"`
	Type           string   `json:"type"`
	Remediation    string   `json:"remediation"`
	TestInfo       []string `json:"test_info"`
	Status         string   `json:"status"`
	ActualValue    string   `json:"actual_value"`
	Scored         bool     `json:"scored"`
	IsMultiple     bool     `json:"IsMultiple"`
	ExpectedResult string   `json:"expected_result"`
	Reason         string   `json:"reason"`
}

func readFile() (*Kubebench, error) {
	jsonFile, err := os.Open(resultFile)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var kb Kubebench

	err = json.Unmarshal(byteValue, &kb)
	if err != nil {
		fmt.Printf("Error in reading kubebench result json file: %v\n", err)
		return nil, err
	}
	id := primitive.NewObjectID()
	t := time.Now()
	fmtS := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	kb.ID = id
	kb.CreateTime = fmtS

	return &kb, nil
}

func create(kubebench *Kubebench, s *Service) (insertedID interface{}, err error) {

	_, err = s.kubebenchCollection.InsertOne(context.Background(), *kubebench)
	if err != nil {
		fmt.Printf("Error in creating kubebench result: %v\n", err)
		return "-1", err
	}

	fmt.Printf("Successful create; Inserted ID: %s\n",kubebench.ID)
	return kubebench.ID, nil

}

func read(id string, s *Service) (kubebench *Kubebench, err error) {
	// convert string to primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(id)

	err = s.kubebenchCollection.FindOne(context.Background(), bson.D{{Key: "_id", Value: oid}}).Decode(&kubebench)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successful read id: %s\n", id)
	return
}

// GetKubebench retrieve a single Kubebench object with "_id" = "id"
// Response: result of the given id
func (s *Service) GetKubebench(c *gin.Context) {
	// query parameters
	id := c.Query("id")

	var kbbench *Kubebench
	kbbench, err := read(id, s)
	if err != nil {
		fmt.Printf("Error in GetKubebench, id=%s: %v\n", id, err)
		util.ResponseError(c, err)
		return
	} else {
		util.ResponseSuccess(c, kbbench, "kubebench")
		return
	}
}

func (s *Service) ListKubebench(c *gin.Context) {
	var kbs []Kubebench
	cursor, err := s.kubebenchCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("Error in list Kubebench results: %v\n", err)
		util.ResponseError(c, err)
		return
	}
	if err = cursor.All(context.TODO(), &kbs); err != nil {
		fmt.Printf("Error in decoding Kubebench result list: %v\n", err)
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, kbs, "kubebench")
}

// NewKubebench run the kubebench script, read from result file, and create result object in DB
// Response: id of the result object
func (s *Service) NewKubebench(c *gin.Context) {

	var kbbench *Kubebench
	var id interface{}
	var err error

	cmd := exec.Command(runScript)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error in waiting for kubebench script executing: %v\n", err)
		goto responseError
	}

	// read from result file
	kbbench, err = readFile()
	if err != nil {
		fmt.Printf("Error in read kubebench result file: %v\n", err)
		goto responseError
	}
	// create result in DB
	id, err = create(kbbench, s)
	if err != nil {
		fmt.Printf("Error in create kubebench result in DB: %v\n", err)
		goto responseError
	}
	util.ResponseSuccess(c, id, "kubebench")
	return

responseError:
	util.ResponseError(c, err)
	return

}

func (s *Service) DeleteKubebench(c *gin.Context) {
	var tmp Kubebench
	c.ShouldBindJSON(&tmp)

	res, err := s.kubebenchCollection.DeleteOne(context.TODO(), bson.M{"id": tmp.ID})
	if err != nil {
		util.ResponseError(c, err)
		return
	}
	util.ResponseSuccess(c, int(res.DeletedCount), "kubebench")
}
