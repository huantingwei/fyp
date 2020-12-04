package kubescore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	// "github.com/zegl/kube-score/scorecard"
	"github.com/gin-gonic/gin"
	db "github.com/huantingwei/fyp/database"
	"go.mongodb.org/mongo-driver/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	dbName = "fyp"
	coll   = "score"
	// windows
	resultFile = "kubescore\\result.json"
	scriptFile = "kubescore\\kubescore.sh"

// linux
// resultFile = "/home/justbadcodes/fyp/kubescore/result.json"
// scriptFile = "/home/justbadcodes/fyp/kubescore/kubescore.sh"
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
	ScoredObjects []ScoredObject `json:"kubescore"`
}

func runScript() {
	// temp name for dev
	namespace := "fyp"

	fmt.Printf("Run kubescore validation...\n")
	cmd := exec.Command("./"+scriptFile, namespace)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error in running kubescore: %v\n", err)
	}
}

func readResultFile() *KubeScore {
	jsonFile, err := os.Open(resultFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var kbscores []ScoredObject

	err = json.Unmarshal(byteValue, &kbscores)
	if err != nil {
		fmt.Printf("Error in reading kubescore result json file: %v\n", err)
	}

	return &KubeScore{kbscores}
}

func createResult(kubescore *KubeScore) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(dbName).Collection(coll)

	res, err := collection.InsertOne(ctx, *kubescore)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Inserted ID: %v\n", res.InsertedID)

}

func ReadWriteFile() {
	fmt.Printf("Start reading kubescore result...\n")
	kbscores := readResultFile()

	fmt.Printf("Start writing into database...\n")
	createResult(kbscores)

	fmt.Printf("Finished writing...\n")
}

func GetAllKubeScore() (kubescores []KubeScore, err error) {
	client, ctx, cancel := db.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(dbName).Collection(coll)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error in reading kubescore: %v\n", err)
		return
	}

	if err = cursor.All(ctx, &kubescores); err != nil {
		return nil, err
	}

	return
	// for _, obj := range (*kbscore).ScoredObjects {
	// 	fmt.Printf("Name: %v\n\n", obj.ObjectName)
	// 	for i, check := range obj.Checks {
	// 		fmt.Printf("Check %v: %v\n", i, check.Check.Name)
	// 		for j, com := range check.Comments {
	// 			fmt.Printf("\tComment %v: %v\n", j, com)
	// 		}
	// 	}
	// }
}

func Print() {
	kbscores, err := GetAllKubeScore()
	if err != nil {
		fmt.Printf("Error in getting all kubescore: %v\n", err)
		return
	}

	kbscore := kbscores[0]

	for _, obj := range kbscore.ScoredObjects {
		fmt.Printf("Name: %v\n\n", obj.ObjectName)
		for i, check := range obj.Checks {
			fmt.Printf("Check %v: %v\n", i, check.Check.Name)
			for j, com := range check.Comments {
				fmt.Printf("\tComment %v: %v\n", j, com)
			}
		}
	}
}

func (s *Service) GetKubescoreResult(c *gin.Context) {
	kbscores, err := GetAllKubeScore()
	if err != nil {
		fmt.Printf("Error in GetKubescoreResult: %v\n", err)
	}
	if len(kbscores) == 0 {
		ReadWriteFile()
		kbscores, _ = GetAllKubeScore()
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"type":  "kubescore",
		"data":  kbscores,
		"count": len(kbscores),
	})
}
