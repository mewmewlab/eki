package utils

import "github.com/docker/docker/client"

var DockerClient *client.Client

func InitDocker() {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	DockerClient = client
}

func CloseDockerClient() {
	DockerClient.Close()
}
