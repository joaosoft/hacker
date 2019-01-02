package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := CmdDependencyGet

	args := os.Args
	if len(args) > 1 {
		cmd = CmdDependency(args[1])
	}

	d, err := NewDependency()
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	switch cmd {
	case CmdDependencyGet:
		if err := d.Get(); err != nil {
			panic(err)
			os.Exit(1)
		}
	case CmdDependencyUpdate:
		if err := d.Update(); err != nil {
			panic(err)
			os.Exit(1)
		}
	case CmdDependencyReset:
		if err := d.Reset(); err != nil {
			panic(err)
			os.Exit(1)
		}
	case CmdDependencyAdd:
		var newImport string
		if len(args) > 2 {
			newImport = args[2]
		}

		if err := d.Add(newImport); err != nil {
			panic(err)
			os.Exit(1)
		}
	case CmdDependencyRemove:
		var removeImport string
		if len(args) > 2 {
			removeImport = args[2]
		}
		if err := d.Remove(removeImport); err != nil {
			panic(err)
			os.Exit(1)
		}
	default:
		fmt.Printf("invalid command! available commands are [%s, %s, %s]", CmdDependencyGet, CmdDependencyUpdate, CmdDependencyReset)
		os.Exit(1)
	}

	os.Exit(0)
}
