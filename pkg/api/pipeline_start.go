package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func PipelineStart(c *core.RequestEvent) error {
	jobID := c.Request.PathValue("id")
	jobs := c.App.Cron().Jobs()
	for _, job := range jobs {
		if job.Id() == jobID {
			go job.Run()
			return c.String(http.StatusOK, "start")
		}
	}
	return c.BadRequestError("job not found", nil)
}
