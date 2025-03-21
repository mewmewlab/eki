package router

import (
	"github.com/mewmewlab/eki/pkg/api"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
)

func RegisterRouter(app *pocketbase.PocketBase) {
	app.OnServe().Bind(&hook.Handler[*core.ServeEvent]{
		Func: func(event *core.ServeEvent) error {

			apiv1 := event.Router.Group("/api/v1")
			{
				docker := apiv1.Group("/docker")
				{
					docker.GET("/{id}/log", api.DockerGetContainerLog)
				}
				pipeline := apiv1.Group("/pipeline")
				{
					pipeline.POST("/create", api.PipelineCreate)
					pipeline.POST("/{id}/run", api.PipelineRun)
					pipeline.GET("/{id}/info", api.PipelineGetInfo)
				}
			}

			return event.Next()
		},
		Priority: 1,
	})
}
