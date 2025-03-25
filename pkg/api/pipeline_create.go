package api

import (
	"github.com/mewmewlab/eki/pkg/utils"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/pocketbase/pocketbase/core"
)

func PipelineCreate(c *core.RequestEvent) error {
	utils.DockerClient.ContainerCreate(
		c.Request.Context(),
		nil, nil, nil, &v1.Platform{}, "")
	return nil
}
