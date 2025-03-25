package modules

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TaskHash(task *Task) string {
	return task.Id
}

type Task struct {
	Id string

	Unit *Unit

	Retry       int
	WaitTimeout int
}

func NewTask(image string) *Task {
	id, _ := gonanoid.New()
	return &Task{
		Id: id,
		Unit: &Unit{
			Image: image,
		},
	}
}

func (t *Task) String() string {
	return t.Id
}
