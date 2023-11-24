package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 构造命令行
var rootCmd = &cobra.Command{
	Use:   "Tool",
	Short: "my tools",
}

// 绑定参数
func init() {
	rootCmd.AddCommand(hashCommand())
	rootCmd.AddCommand(encryptCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
