package api

import (
	"net/http"
	"slices"
	"strings"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
)

func PipelineGet(c *core.RequestEvent) error {
	jobs := c.App.Cron().Jobs()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"jobs": slices.DeleteFunc(jobs, func(job *cron.Job) bool {
			return strings.HasPrefix(job.Id(), "__pb")
		}),
	})
}

func PipelineGetInfo(c *core.RequestEvent) error {
	return c.String(http.StatusOK, c.Request.PathValue("id"))
}
