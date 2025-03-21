package api

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pocketbase/pocketbase/core"
)

type DockerCreateContainerRequest struct {
	Image string `json:"image"`
}

type DockerCreateContainerResponse struct {
}

func DockerCreateContainer(c *core.RequestEvent) error {
	dockerClient.ContainerCreate(
		c.Request.Context(),
		nil, nil, nil, &v1.Platform{}, "")
	return nil
}
