package cmd

import (
	"os"

	"github.com/latiif/lail/cmd/repl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lail",
	Short: "Interpreter for lail ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			repl.Start(os.Stdin, os.Stdout)
		} else {
			for _, file := range args {
				fileHandle, err := os.Open(file)
				if err != nil {
					continue
				}
				repl.InterpretFile(fileHandle, os.Stdout)
				fileHandle.Close()
			}
		}
	},
}

// Executes the program
func Execute() error {
	return rootCmd.Execute()
}
