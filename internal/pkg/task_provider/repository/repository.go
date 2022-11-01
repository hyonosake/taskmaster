package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
	"go.uber.org/zap"

	"github.com/hyonosake/taskmaster/internal/pkg/task_provider/entity"
)

type taskRepository struct {
	logger *zap.SugaredLogger
}

func NewTaskRepository(log *zap.SugaredLogger) *taskRepository {
	return &taskRepository{logger: log}
}

func (t *taskRepository) GetTasks(req entity.GetTasksRequest) (cfg entity.Config, err error) {

	absPath, err := filepath.Abs(req.Path)
	t.logger.Debugf("check GetTasks file : %s", absPath)

	if err != nil {
		return cfg, fmt.Errorf("path %s in not absolute", req.Path)
	}
	f, err := os.Open(absPath)
	if err != nil {
		return cfg, fmt.Errorf("unable to open %s: %w", req.Path, err)
	}
	if f == nil {
		return cfg, fmt.Errorf("unable to open %s", req.Path)
	}
	defer func() {
		if err = f.Close(); err != nil {
			err = fmt.Errorf("unable to close file: %w", err)
		}
	}()
	cfg, err = t.convertFromRawData(f)
	return cfg, nil
}

func (t *taskRepository) convertFromRawData(reader io.Reader) (entity.Config, error) {
	var cfg entity.Config

	rawBytes, err := io.ReadAll(reader)
	if err != nil {
		return entity.Config{}, fmt.Errorf("unable to read data: %w", err)
	}

	err = yaml.Unmarshal(rawBytes, &cfg)
	if err != nil {
		return entity.Config{}, fmt.Errorf("unable to parse data: %w", err)
	}

	return cfg, err
}
