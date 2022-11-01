package main

import (
	"log"

	"github.com/hyonosake/taskmaster/internal/pkg/app"
	task_repository "github.com/hyonosake/taskmaster/internal/pkg/task_provider/repository"
	"github.com/hyonosake/taskmaster/internal/pkg/taskmaster/usecase"
	"github.com/hyonosake/taskmaster/internal/pkg/utils"
)

func main() {

	cfg, err := app.NewConfig()
	if err != nil {
		log.Fatalf("unable to run taskmaster: %v", err)
	}

	logger, _ := utils.NewSugaredLogger(cfg.GetLogLevel(), cfg.GetLogFile())
	taskProvider := task_repository.NewTaskRepository(logger)

	master := usecase.NewTaskmaster(taskProvider)
	tasks, err := master.GetTasks(cfg.GetConfigPath())
	_, _ = err, tasks
	logger.Debug("provided config", tasks)
}
