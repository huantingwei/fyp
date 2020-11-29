package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// cmd := exec.Command("kubectl", "get", "all", "-o", "json")
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// var pod struct {
	// 	APIVersion string
	// 	Items      []interface{}
	// 	Kind       string
	// 	Metadata   interface{}
	// 	// Spec       string
	// }
	// if err := json.NewDecoder(stdout).Decode(&pod); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s, %s, %s, %s\n", pod.APIVersion, pod.Items, pod.Kind, pod.Metadata)

	// pod := script.GetPod()
	// fmt.Println(pod)

	cmd := exec.Command("kubectl", "get", "pods", "-o=jsonpath='{.items[0]}'")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("type: %T\n", out)
	fmt.Printf("combined out:\n%s\n", string(out))
}
