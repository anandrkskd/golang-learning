package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func main() {
	mesheryURL := "https://api.github.com/repos/layer5io/meshery/releases/latest" // "https://api.github.com/repos/layer5io/meshery/releases"
	resp, err := http.Get(mesheryURL)
	downloadMesheryURL := "https://github.com/layer5io/meshery/releases/download"
	if err != nil {
		// download the default version as 1.24.1 if unable to fetch latest page data
		fmt.Print("error")
	} else {
		var dat map[string]interface{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(body, &dat); err != nil {
			log.Fatal(err)
		}
		num := dat["tag_name"]
		os := ""
		arch := ""

		if runtime.GOOS == "windows" {
			// Meshery running on Windows host
			os = "Windows"
		} else if runtime.GOOS == "linux" {
			// Meshery running on Linux host
			os = "Linux"
		} else {
			// Assume Meshery running on MacOS host
			os = "Darwin"
		}

		if runtime.GOARCH == "amd64" {
			// Meshery running on x86_64 host
			arch = "x86_64"
		} else {
			// Meshery running on i386 host
			arch = "i386"
		}

		version := fmt.Sprint("%v", num)
		downloadMesheryURL = fmt.Sprintf(downloadMesheryURL+"/%v/mesheryctl_%v_%v_%v.zip", num, version[3:], os, arch)
	}
	//osdetails := "mac"
	//dockerComposeBinaryURL = dockerComposeBinaryURL + "-" + osdetails
	fmt.Println(downloadMesheryURL)

}
