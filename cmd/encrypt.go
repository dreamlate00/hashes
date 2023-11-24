package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func encryptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Perform symmetric encryption",
		Run: func(cmd *cobra.Command, args []string) {
			// 处理对称加密的逻辑
			fmt.Println("Performing encryption...")
		},
	}
	cmd.Flags().StringVarP(&hashType, "type", "t", "AES", "encrypt type (AES)")
	return cmd
}
