package modules

import (
	"github.com/dominikbraun/graph"
)

type Pipeline struct {
	dag     graph.Graph[string, *Task]
	TaskIDs []string
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		dag:     graph.New(TaskHash, graph.Directed(), graph.PreventCycles()),
		TaskIDs: nil,
	}
}

func (p *Pipeline) AddTask(t *Task) error {
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

func (p *Pipeline) GetTask(id string) (*Task, error) {
	return p.dag.Vertex(id)
}
