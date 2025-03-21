package pipeline_test

import (
	"testing"

	"github.com/mewmewlab/eki/pkg/modules/pipeline"
	"github.com/mewmewlab/eki/pkg/modules/task"
)

func TestPipeline(t *testing.T) {
	p := pipeline.NewPipeline()
	a := task.NewTask("A")
	b := task.NewTask("B")
	c := task.NewTask("C")
	d := task.NewTask("D")
	p.AddTask(a)
	p.AddTask(b)
	p.AddTask(c)
	p.AddTask(d)
	p.AddDependency(a.String(), b.String())
	p.AddDependency(a.String(), c.String())
	p.AddDependency(b.String(), c.String())
	p.AddDependency(b.String(), d.String())
	p.AddDependency(c.String(), d.String())
	if err := p.TopologicalSort(); err != nil {
		t.Error(err)
	}
	for _, id := range p.TaskIDs {
		task, err := p.GetTask(id)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(task.Image)
	}
}
