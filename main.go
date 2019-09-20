package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joaosoft/logger"
)

func main() {
	hacker, err := NewHacker()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	args := os.Args
	var cmd Command
	if len(args) > 1 {
		cmd = Command(args[1])
	}

	switch cmd {
	case CommandRun:
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
		logger.Info("started!")
		for {
			select {
			case <-termChan:
				logger.Info("received term signal")
				hacker.pm.Stop()
				os.Exit(0)

			case <-hacker.quit:
				logger.Info("received shutdown signal")
				hacker.pm.Stop()
				os.Exit(0)

			case <-time.After(time.Second * 10):
				logger.Info("running in background!")
			}
		}
	default:
		logger.Error("unknown command!")
	}

	os.Exit(0)
}
