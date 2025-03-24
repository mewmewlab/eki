package api

import (
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/pocketbase/pocketbase/core"
)

func DockerGetContainerLog(c *core.RequestEvent) error {
	id := c.Request.PathValue("id")
	c.App.Logger().Info(id)
	stream, err := dockerClient.ContainerLogs(c.Request.Context(), id, container.LogsOptions{
		Follow:     true,
		Timestamps: true,
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return c.Error(http.StatusBadRequest, "container id not found", err)
	}
	defer stream.Close()
	return c.Stream(http.StatusOK, "text/plain", stream)
}
