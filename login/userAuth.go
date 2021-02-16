package login

import (
	"fmt"
	"io"

	//"io/ioutil"
	"bufio"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/huantingwei/fyp/util"
)

type Service struct {
	cmdStdinPipe io.WriteCloser
}

type Project struct {
	ClusterName string `json:"clusterName"`
	ProjectName string `json:"projectName"`
	ZoneName    string `json:"zoneName"`
}

type GoogleCode struct {
	Code string `json:"code"`
}

func NewService(r *gin.RouterGroup) {
	s := &Service{}

	go s.asyncBash()

	r.POST("/login", s.postProjectInfo)
	r.POST("/googleAuth", s.postGoogleCode)

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

func (s *Service) postProjectInfo(c *gin.Context) {

	var p Project
	c.ShouldBindJSON(&p)

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
	}
	util.ResponseSuccess(c, verificationCode, "verification")
}
