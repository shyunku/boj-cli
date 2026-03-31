package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"boj/internal/boj"
	"boj/internal/runner"
	"boj/internal/workspace"

	"github.com/spf13/cobra"
)

var testLanguage string

var testCmd = &cobra.Command{
	Use:   "test <problem_id> [file]",
	Short: "Test a solution against example test cases",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		var file string
		if len(args) == 2 {
			file = args[1]
		} else {
			var err error
			file, err = pickFile(id, testLanguage)
			if err != nil {
				return err
			}
		}

		problem, err := boj.GetProblem(id)
		if err != nil {
			return err
		}

		if len(problem.TestCases) == 0 {
			return fmt.Errorf("no test cases found for problem %s", id)
		}

		lang := testLanguage
		if lang == "" {
			lang = workspace.EnvFromFile(file)
		}
		results, err := runner.RunTests(file, lang, problem.TestCases)
		if err != nil {
			return err
		}

		passed := 0
		for i, r := range results {
			if r.Passed {
				passed++
				fmt.Printf("\033[32m[PASS]\033[0m Test case %d\n", i+1)
			} else {
				fmt.Printf("\033[31m[FAIL]\033[0m Test case %d\n", i+1)
				fmt.Printf("  Expected: %q\n", r.Expected)
				fmt.Printf("  Got:      %q\n", r.Actual)
			}
		}

		fmt.Printf("\n%d/%d passed\n", passed, len(results))
		return nil
	},
}

func pickFile(id, lang string) (string, error) {
	// If language is specified, infer the file directly
	if lang != "" {
		ext, err := workspace.ExtForEnv(lang)
		if err != nil {
			return "", err
		}
		file := fmt.Sprintf("%s.%s", id, ext)
		if _, err := os.Stat(file); err != nil {
			return "", fmt.Errorf("file %s not found\nHint: use `boj code %s -l %s` to create one", file, id, lang)
		}
		return file, nil
	}

	files, err := workspace.FindFiles(id)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", fmt.Errorf("no solution files found for problem %s\nHint: use `boj code %s -l nodejs` to create one", id, id)
	}
	if len(files) == 1 {
		return files[0], nil
	}

	fmt.Printf("Multiple files found for problem %s:\n\n", id)
	for i, f := range files {
		fmt.Printf("  [%d] %s\n", i+1, f)
	}
	fmt.Printf("\nSelect a file (1-%d): ", len(files))

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil || n < 1 || n > len(files) {
		return "", fmt.Errorf("invalid selection")
	}
	return files[n-1], nil
}

func init() {
	testCmd.Flags().StringVarP(&testLanguage, "language", "l", "", "Language/runtime to infer file (e.g. nodejs → 1000.js)")
}
