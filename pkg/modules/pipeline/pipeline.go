package pipeline

import (
	"github.com/dominikbraun/graph"
	"github.com/mewmewlab/eki/pkg/modules/task"
)

type Pipeline struct {
	dag     graph.Graph[string, *task.Task]
	TaskIDs []string
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		dag:     graph.New(task.TaskHash, graph.Directed(), graph.PreventCycles()),
		TaskIDs: nil,
	}
}

func (p *Pipeline) AddTask(t *task.Task) error {
	return p.dag.AddVertex(t)
}

func (p *Pipeline) AddDependency(from, to string) error {
	return p.dag.AddEdge(from, to)
}

func (p *Pipeline) TopologicalSort() error {
	tasks, err := graph.TopologicalSort(p.dag)
	if err != nil {
		return err
	}
	p.TaskIDs = tasks
	return nil
}

func (p *Pipeline) GetTask(id string) (*task.Task, error) {
	return p.dag.Vertex(id)
}
