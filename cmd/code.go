package cmd

import (
	"fmt"

	"boj/internal/boj"
	"boj/internal/workspace"

	"github.com/spf13/cobra"
)

var codeLanguage string

var codeCmd = &cobra.Command{
	Use:   "code <problem_id>",
	Short: "Create a solution file for a problem",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if codeLanguage == "" {
			return fmt.Errorf("language is required: use -l or --language (see `boj langs`)")
		}

		id := args[0]

		problem, err := boj.GetProblem(id)
		if err != nil {
			return err
		}

		path, err := workspace.CreateFile(problem, codeLanguage)
		if err != nil {
			return err
		}

		fmt.Printf("Created: %s\n", path)
		return nil
	},
}

func init() {
	codeCmd.Flags().StringVarP(&codeLanguage, "language", "l", "", "Language/runtime to use (see `boj langs`)")
}
