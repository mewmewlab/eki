package modules_test

import (
	"testing"

	"github.com/mewmewlab/eki/pkg/modules"
)

func TestNewTask(t *testing.T) {
	for i := 0; i < 10; i++ {
		task := modules.NewTask("1")
		if task.Id == "" {
			t.Errorf("Expected task ID to be set, got empty string")
		}
		t.Log("Task ID:", task.Id)
	}
}
