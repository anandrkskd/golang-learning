package main

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func check(s []string, atempt int) {
	count := 0
	for _, x := range s {
		if strings.TrimRight(x, " ") == "meshery_meshery_1" {
			count++
		}
	}
	if count < 1 && atempt < 1 {
		//fmt.Println("service cannot be started, attempting to restart the service!")
		exec.Command("mesheryctl", "stop")
		exec.Command("mesheryctl", "start")
		atempt++
		check(s, atempt)
	} else if atempt >= 1 {
		//fmt.Println("service cannot be started. Showing meshery logs")
		exec.Command("mesheryctl", "logs")
	} else {

	}
}
func main() {
	exec.Command("echo", "$")
	exec.Command("meshery", "start")
	//container := [5]string{"meshery_1", "meshery_meshery-consul_1", "meshery_watchtower_1", "meshery_meshery-istio_1", "meshery_meshery-linkerd_1"}
	cmd := exec.Command("docker", "ps", "--format", "{{.Names}}")
	out, err := cmd.Output()

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"number": 0,
	}).Trace("Went to the beach")
	//log.Error(out)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	s := strings.Split(string(out), "\n")
	check(s, 0)
}
