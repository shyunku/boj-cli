package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boj",
	Short: "Baekjoon Online Judge CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(codeCmd)
	rootCmd.AddCommand(resumeCmd)
	rootCmd.AddCommand(langsCmd)
	rootCmd.AddCommand(searchCmd)
}
