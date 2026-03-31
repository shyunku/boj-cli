package cmd

import (
	"fmt"

	"boj/internal/boj"
	"boj/internal/workspace"

	"github.com/spf13/cobra"
)

var codeEnvironment string

var codeCmd = &cobra.Command{
	Use:   "code <problem_id>",
	Short: "Create a solution file for a problem",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		problem, err := boj.GetProblem(id)
		if err != nil {
			return err
		}

		path, err := workspace.CreateFile(problem, codeEnvironment)
		if err != nil {
			return err
		}

		fmt.Printf("Created: %s\n", path)
		return nil
	},
}

func init() {
	codeCmd.Flags().StringVarP(&codeEnvironment, "environment", "e", "nodejs", "Runtime environment (nodejs, python, cpp)")
}
