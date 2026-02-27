package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "awesome-docker",
		Short: "Quality tooling for the awesome-docker curated list",
	}
	root.AddCommand(
		&cobra.Command{Use: "version", Short: "Print version", Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("awesome-docker v0.1.0")
		}},
	)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
