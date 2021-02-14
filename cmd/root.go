package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/latiif/lail/cmd/repl"
)

func execute(args []string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Fatal issue detected:", r, "\nABORTED!")
		}
	}()
	if len(args) == 0 {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		for _, file := range args {
			fileHandle, err := os.Open(file)
			if err != nil {
				continue
			}
			repl.InterpretFile(filepath.Dir(file), fileHandle, os.Stdout, os.Stderr)
			fileHandle.Close()
		}
	}
	return nil
}

// Executes the program
func Execute() error {
	return execute(os.Args[1:])
}
