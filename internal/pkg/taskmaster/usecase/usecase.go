package usecase

import (
	"github.com/hyonosake/taskmaster/internal/pkg/task_provider"
	"github.com/hyonosake/taskmaster/internal/pkg/task_provider/entity"
)

type Taskmaster struct {
	taskProvider task_provider.TaskProvider
}

func NewTaskmaster(taskProvider task_provider.TaskProvider) *Taskmaster {
	return &Taskmaster{
		taskProvider: taskProvider,
	}
}

func (t *Taskmaster) GetTasks(path string) (entity.Config, error) {
	req := entity.GetTasksRequest{Path: path}
	return t.taskProvider.GetTasks(req)
}
