package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type settings struct {
	ConfigPath string `yaml:"config_path"`
	LogPath    string `yaml:"log_path"`
	LogLevel   uint8  `yaml:"log_level"`
}

type initCfg struct {
	Settings settings `json:"settings"`
}

func NewConfig() (Config, error) {

	var cfg Config
	rawCfg, err := getRawConfig()
	if err != nil {
		return cfg, err
	}

	cfg.ConfigPath = rawCfg.Settings.ConfigPath
	if rawCfg.Settings.LogPath != "" {
		cfg.LogFile, err = os.OpenFile(rawCfg.Settings.LogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return cfg, fmt.Errorf("unable to open %s: %w", rawCfg.Settings.ConfigPath, err)
		}
	}
	return cfg, nil
}

func getRawConfig() (initCfg, error) {

	path := os.Getenv("TASKMASTER_CFG_PATH")
	absPath, err := filepath.Abs(path)
	if err != nil {
		return initCfg{}, fmt.Errorf("path %s in not absolute", path)
	}

	log.Printf("check file : %s", absPath)
	f, err := os.Open(absPath)
	if err != nil {
		return initCfg{}, fmt.Errorf("unable to open %s: %w", path, err)
	}
	if f == nil {
		return initCfg{}, fmt.Errorf("unable to open %s", path)
	}
	defer func() {
		if err = f.Close(); err != nil {
			err = fmt.Errorf("unable to close file: %w", err)
		}
	}()
	rawBytes, err := io.ReadAll(f)
	if err != nil {
		return initCfg{}, fmt.Errorf("unable to read data: %w", err)
	}
	cfg := initCfg{}
	err = yaml.Unmarshal(rawBytes, &cfg)
	if err != nil {
		return initCfg{}, fmt.Errorf("unable to parse data: %w", err)
	}
	return cfg, nil
}
