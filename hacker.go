package main

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Hacker struct {
	config        *HackerConfig
	quit          chan int
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
}

func NewHacker(options ...HackerOption) (*Hacker, error) {
	config, simpleConfig, err := NewConfig()
	log := logger.NewLogDefault("hacker", logger.WarnLevel)

	service := &Hacker{
		quit:   make(chan int),
		pm:     manager.NewManager(manager.WithRunInBackground(true), manager.WithLogLevel(logger.WarnLevel)),
		logger: log,
		config: config.Hacker,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		log.Error(err.Error())
	} else if config.Hacker != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Hacker.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.Reconfigure(options...)

	return service, nil
}

func (d *Hacker) Run() error {
	d.logger.Debug("executing Run")
	return nil
}
