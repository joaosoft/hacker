package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joaosoft/logger"
)

func main() {
	cmd := CmdDependencyGet
	var err error

	args := os.Args
	if len(args) > 1 {
		cmd = CmdDependency(args[1])
	}

	d, err := NewDependency()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	// listen for termination signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		for {
			select {
			case <-termChan:
				logger.Info("received term signal")
				d.doUndoBackupVendor()
				os.Exit(0)

			case <-d.quit:
				logger.Info("received shutdown signal")
				d.doUndoBackupVendor()
				os.Exit(0)
			}
		}
	}()

	switch cmd {
	case CmdDependencyGet:
		err = d.Get()
	case CmdDependencyUpdate:
		if err := d.Update(); err != nil {
			panic(err)
			os.Exit(1)
		}
	case CmdDependencyReset:
		err = d.Reset()
	case CmdDependencyAdd:
		var newImport string
		if len(args) > 2 {
			newImport = args[2]
		}

		err = d.Add(newImport)
	case CmdDependencyRemove:
		var removeImport string
		if len(args) > 2 {
			removeImport = args[2]
		}
		err = d.Remove(removeImport)
	default:
		fmt.Printf("invalid command! available commands are [%s, %s, %s]", CmdDependencyGet, CmdDependencyUpdate, CmdDependencyReset)
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
