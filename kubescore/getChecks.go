package kubescore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Check struct {
	CheckMeta map[string]interface{}
	Grade     int
	Skipped   bool
	Comments  interface{}
}
type KubeScore struct {
	ObjectName string
	TypeMeta   map[string]string
	ObjectMeta map[string]string
	Checks     []Check
}

const fileName = "result.json"

func readFile() *[]KubeScore {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var kbscores []KubeScore

	err = json.Unmarshal(byteValue, &kbscores)
	if err != nil {
		fmt.Println(err)
	}

	return &kbscores
}

func GetScore() {
	kbscores := readFile()

	for _, obj := range *kbscores {
		fmt.Printf("Name: %v\n", obj.ObjectName)
		for i, check := range obj.Checks {
			fmt.Printf("Check %v: %v\n", i, check.Comments)
		}
	}
}
