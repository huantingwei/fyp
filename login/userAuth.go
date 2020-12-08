package login

import (
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"os"
	"os/exec"
	"net/http"
	"bufio"

	"github.com/gin-gonic/gin"
)

type Service struct{
	cmdStdinPipe			io.WriteCloser
}

func NewService(r *gin.RouterGroup){
	s := &Service{};

	go s.asyncBash()

	r.POST("/", s.postProjectInfo);
	r.POST("/googleAuth", s.postGoogleCode)

}

func (s *Service) asyncBash(){
	cmd := exec.Command("./login/kubebench.sh")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	s.cmdStdinPipe = stdin

	cmd.Run()
}

func (s *Service) postProjectInfo(c *gin.Context){
	clusterName := c.PostForm("clusterName");
	projectName := c.PostForm("projectName");
	zoneName := c.PostForm("zoneName");

	io.WriteString(s.cmdStdinPipe, clusterName + "\n")
	io.WriteString(s.cmdStdinPipe, projectName + "\n")
	io.WriteString(s.cmdStdinPipe, zoneName + "\n")

	urlChan := make(chan string)

	go func(){
		urlChan <- readURLFromFile()
	}()

	url := <- urlChan
	c.IndentedJSON(http.StatusOK, gin.H{
		"URL": url,
	})
}

func readURLFromFile() string{
	for{
		if f, err := os.Stat("./url.txt") ; err == nil {
			if f.Size() > 0{
				file, err2 := os.Open("./url.txt")
				if err2 != nil {
					fmt.Printf(err2.Error());
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				scanner.Scan();
				return scanner.Text()
			}
		}
	}
}

func (s *Service) postGoogleCode(c *gin.Context){
	verificationCode := c.PostForm("code")
	//io.WriteString(s.cmdStdinPipe, verificationCode + "\n")
	f, _ := os.Create("./token.txt")
	_, err := f.WriteString(verificationCode)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code: %s\n", verificationCode)
	c.IndentedJSON(http.StatusOK, gin.H{
		"code": verificationCode,
	})
}