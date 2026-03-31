package cmd

import (
	"fmt"

	"boj/internal/workspace"

	"github.com/spf13/cobra"
)

var langsCmd = &cobra.Command{
	Use:   "langs",
	Short: "List available environments",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		for _, e := range workspace.Environments() {
			fmt.Printf("  %-10s %s\n", e.Name, e.Ext)
		}
	},
}
