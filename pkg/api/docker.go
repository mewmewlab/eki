package api

import (
	"github.com/docker/docker/client"
)

var (
	dockerClient *client.Client
)

func initDocker() {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	dockerClient = client
}

func CloseDockerClient() {
	dockerClient.Close()
}
