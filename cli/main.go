package main

import (
	"fmt"
	"os"

	commands "gobasic/cli/command"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "A basic cli with Cobra!",
	}

	rootCmd.AddCommand(
		commands.WhereAmI(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
