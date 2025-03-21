package api

import (
	"bufio"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/subscriptions"
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
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		data := scanner.Bytes()
		c.App.Logger().Info(string(data))
		jsonData, _ := json.Marshal(map[string]interface{}{
			"data": string(data),
		})
		clients := c.App.SubscriptionsBroker().Clients()
		for _, client := range clients {
			client.Send(subscriptions.Message{
				Name: "docker.container.log." + id,
				Data: jsonData,
			})
		}
	}
	return nil
}
