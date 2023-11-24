package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show manager-api version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("1.0.0")
		},
	}
}
