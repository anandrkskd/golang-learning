package main

import (

	//"log"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	cmd := exec.Command("ls", "/home/leo")
	out, _ := cmd.Output()
	a := strings.Split(string(out), "\n")
	log.Print(a)
}
