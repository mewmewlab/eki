package task

import (
	"github.com/docker/docker/api/types/container"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TaskHash(task *Task) string {
	return task.Id
}

type Task struct {
	Id string

	Image           string
	ContainerName   string
	ContainerConfig *container.Config
	HostConfig      *container.HostConfig

	Retry       int
	WaitTimeout int
}

func NewTask(image string) *Task {
	id, _ := gonanoid.New()
	return &Task{
		Id:    id,
		Image: image,
	}
}

func (t *Task) String() string {
	return t.Id
}
