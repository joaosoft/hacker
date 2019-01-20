package main

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Dependency *DependencyConfig `json:"dependency"`
}

// DependencyConfig ...
type DependencyConfig struct {
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
	Log      struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*DependencyConfig, error) {
	appConfig := &AppConfig{}
	if _, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig); err != nil {
		log.Error(err.Error())

		return &DependencyConfig{}, err
	}

	return appConfig.Dependency, nil
}
