package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// 参数
var (
	hashType string
)

func hashCommand() *cobra.Command {
	hashCmd := &cobra.Command{
		Use:   "hash",
		Short: "Calculate hash value",
		Run: func(cmd *cobra.Command, args []string) {
			// 处理计算哈希值的逻辑
			run(cmd, args)
		},
	}
	hashCmd.Flags().StringVarP(&hashType, "type", "t", "md5", "Hash type (md5 , sha1 , sha256 or base64)")
	return hashCmd
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a string")
		return
	}

	str := strings.Join(args, " ")
	hash := generateHash(str, hashType)
	fmt.Printf("%s Hash: %s\n", strings.ToUpper(hashType), hash)
}

func generateHash(str string, hashType string) string {
	switch hashType {
	case "md5":
		hasher := md5.New()
		hasher.Write([]byte(str))
		return hex.EncodeToString(hasher.Sum(nil))
	case "sha1":
		hasher := sha1.New()
		hasher.Write([]byte(str))
		return hex.EncodeToString(hasher.Sum(nil))
	case "sha256":
		hasher := sha256.New()
		hasher.Write([]byte(str))
		return hex.EncodeToString(hasher.Sum(nil))
	case "base64":
		return base64.StdEncoding.EncodeToString([]byte(str))
	default:
		return ""
	}
}
