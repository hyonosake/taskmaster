package task_provider

import "github.com/hyonosake/taskmaster/internal/pkg/task_provider/entity"

type TaskProvider interface {
	GetTasks(request entity.GetTasksRequest) (entity.Config, error)
	//UpdateTasks() ([]entity.Task, error)
}
