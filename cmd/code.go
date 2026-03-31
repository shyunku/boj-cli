package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

		fmt.Printf("Opening: %s\n", path)
		return openEditor(path)
	},
}

func openEditor(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	codeCmd.Flags().StringVarP(&codeLanguage, "language", "l", "", "Language/runtime to use (see `boj langs`)")
}
