package command

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var personName string

// WhereAmI command-line, it will display CWD:
func WhereAmI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whereami [OPTIONS]",
		Short: "Display absolute path of current workfing directory.",
		Run: func(cmd *cobra.Command, args []string) {
			handle(cmd, args)
		},
		DisableFlagsInUseLine: true,
	}

	cmd.Flags().StringVar(&personName, "name", "there", "Your name")

	return cmd
}

func handle(cmd *cobra.Command, args []string) {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Hello " + personName + "!")
	fmt.Println("You are at: ", cwd)
}
