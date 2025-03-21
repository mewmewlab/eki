package api

import (
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pocketbase/pocketbase/core"
)

func PipelineCreate(c *core.RequestEvent) error {
	dockerClient.ContainerCreate(
		c.Request.Context(),
		nil, nil, nil, &v1.Platform{}, "")
	return nil
}

func PipelineRun(c *core.RequestEvent) error {
	return nil
}
