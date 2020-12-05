package kubebench

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	db "github.com/huantingwei/fyp/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	dbName = "fyp"
	coll   = "kubebench"
	// windows
	resultFile = "kubebench\\result.json"
	scriptFile = "kubebench\\kubebench.sh"

// linux
// resultFile = "/home/justbadcodes/fyp/kubescore/result.json"
// scriptFile = "/home/justbadcodes/fyp/kubescore/kubescore.sh"
)

type Kubebench struct {
	Chapters []Chapter
}

type Chapter struct {
	ID       string    `json:"id"`
	Version  string    `json:"version"`
	Text     string    `json:"text"`
	NodeType string    `json:"node_type"`
	Sections []Section `json:"tests"`

	TotalPass int `json:"total_pass"`
	TotalFail int `json:"total_fail"`
	TotalWarn int `json:"total_warn"`
	TotalInfo int `json:"total_info"`
}

type Section struct {
	Section string   `json:"section"`
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

func readResultFile() (*Kubebench, error) {
	jsonFile, err := os.Open(resultFile)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res []Chapter

	err = json.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Printf("Error in reading kubebench result json file: %v\n", err)
		return nil, err
	}

	return &Kubebench{res}, nil
}

func createResult(kubebench *Kubebench) error {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(dbName).Collection(coll)

	res, err := collection.InsertOne(ctx, *kubebench)
	if err != nil {
		fmt.Printf("Error in creating kubebench result: %v\n", err)
		return err
	}

	fmt.Printf("Inserted ID: %v\n", res.InsertedID)
	return nil

}

func readWriteResult() error {
	fmt.Printf("Start reading kubescore result...\n")
	kubebench, err := readResultFile()
	if err != nil {
		fmt.Printf("Error in readResultFile: %v\n", err)
		return err
	}

	fmt.Printf("Start writing into database...\n")
	err = createResult(kubebench)
	if err != nil {
		fmt.Printf("Error in createResult: %v\n", err)
		return err
	}
	fmt.Printf("Finished writing...\n")
	return nil
}

func GetAllKubebench() (kubebench []Kubebench, err error) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(dbName).Collection(coll)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error in reading kubebench: %v\n", err)
		return nil, err
	}

	if err = cursor.All(ctx, &kubebench); err != nil {
		return nil, err
	}

	return kubebench, nil
}

func (s *Service) GetKubebenchResult(c *gin.Context) {
	kubebench, err := GetAllKubebench()
	if err != nil {
		fmt.Printf("Error in GetKubebenchResult: %v\n", err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"type":  "kubebench",
		"data":  kubebench,
		"count": len(kubebench),
	})
}
