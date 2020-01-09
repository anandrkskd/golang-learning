package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if _, err := os.Stat("/usr/local/bin/docker-compose"); os.IsNotExist(err) {
		log.Print("not present")
	} else {
		log.Print("present")
	}
	if err := exec.Command("docker-compose", "-v").Run(); err == nil {
		log.Print("present")
	} else {
		log.Print("not present")
	}
}
