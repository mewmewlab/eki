package modules_test

import (
	"testing"

	"github.com/mewmewlab/eki/pkg/modules"
)

func TestPipeline(t *testing.T) {
	p := modules.NewPipeline()
	a := modules.NewTask("A")
	b := modules.NewTask("B")
	c := modules.NewTask("C")
	d := modules.NewTask("D")
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
		t.Log(task.Unit.Image)
	}
}
