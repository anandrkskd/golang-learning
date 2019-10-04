//check for latest docker-compose version present
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var dockerComposeBinaryURL string = "https://github.com/docker/compose/releases/download/"

func main() {

	fmt.Println(new())
}

func new() string {
	resp, err := http.Get("https://api.github.com/repos/docker/compose/releases/latest")
	if err != nil {
		// handle error
	}
	var dat map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	num := dat["tag_name"]
	//	fmt.Println(num)
	dockerComposeBinaryURL = fmt.Sprintf(dockerComposeBinaryURL+"%v/docker-compose", num)
	return string(dockerComposeBinaryURL)
}
