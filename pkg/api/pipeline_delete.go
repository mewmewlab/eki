package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func PipelineDelete(c *core.RequestEvent) error {

	jobID := c.Request.PathValue("id")
	jobs := c.App.Cron().Jobs()
	for _, job := range jobs {
		if job.Id() == jobID {
			c.App.Cron().Remove(jobID)
			return c.String(http.StatusOK, "deleted")
		}
	}
	return c.BadRequestError("job not found", nil)
}
