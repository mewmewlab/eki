package api

import (
	"net/http"

	"github.com/docker/docker/api/types/image"
	"github.com/mewmewlab/eki/pkg/utils"
	"github.com/pocketbase/pocketbase/core"
)

type UnitCreateRequest struct {
	Image string `json:"image"`
}

func UnitCreate(c *core.RequestEvent) error {
	req := &UnitCreateRequest{}
	if err := c.BindBody(req); err != nil {
		return c.BadRequestError("create unit failed", err)
	}
	reader, err := utils.DockerClient.ImagePull(c.Request.Context(), req.Image, image.PullOptions{})
	if err != nil {
		return c.BadRequestError("pull image failed", err)
	}
	defer reader.Close()
	return c.Stream(http.StatusOK, "text/plain", reader)
}
