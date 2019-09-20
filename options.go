package main

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// HackerOption ...
type HackerOption func(hacker *Hacker)

// Reconfigure ...
func (d *Hacker) Reconfigure(options ...HackerOption) {
	for _, option := range options {
		option(d)
	}
}

// WithConfiguration ...
func WithConfiguration(config *HackerConfig) HackerOption {
	return func(hacker *Hacker) {
		hacker.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) HackerOption {
	return func(hacker *Hacker) {
		hacker.logger = logger
		hacker.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) HackerOption {
	return func(hacker *Hacker) {
		hacker.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) HackerOption {
	return func(hacker *Hacker) {
		hacker.pm = mgr
	}
}
