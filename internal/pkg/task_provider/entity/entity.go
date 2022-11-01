package entity

import "time"

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name string   `yaml:"name"`
	Opts TaskOpts `yaml:"options"`
}

type TaskOpts struct {
	Cmd       string        `yaml:"command"`
	Flags     string        `yaml:"flags"`
	Stdin     string        `yaml:"stdin"`
	Stdout    string        `yaml:"stdout"`
	Stderr    string        `yaml:"stderr"`
	Repeating bool          `yaml:"repeating"`
	Interval  time.Duration `yaml:"interval"`
}

type GetTasksRequest struct {
	Path string
}

func (c *Config) GetTasks() []Task {
	return c.Tasks
}
