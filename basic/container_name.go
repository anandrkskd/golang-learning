package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var cont []string

func main() {
	//	_ = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	//a := []string{"/meshery_meshery_1"}
	b := "/meshery_meshery_1"
	for _, container := range containers {
		if b == container.Names[0] {
			fmt.Println(container.Names)
		}
		cont = append(cont, container.Names[0])
	}

	fmt.Println(cont)
}
