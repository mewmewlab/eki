package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func PipelineUpdate(c *core.RequestEvent) error {
	return c.String(http.StatusOK, c.Request.PathValue("id"))
}
