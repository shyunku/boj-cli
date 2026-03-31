package cmd

import (
	"fmt"

	"boj/internal/boj"
	"boj/internal/runner"

	"github.com/spf13/cobra"
)

var testEnvironment string

var testCmd = &cobra.Command{
	Use:   "test <problem_id> <file>",
	Short: "Test a solution against example test cases",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		file := args[1]

		problem, err := boj.GetProblem(id)
		if err != nil {
			return err
		}

		if len(problem.TestCases) == 0 {
			return fmt.Errorf("no test cases found for problem %s", id)
		}

		results, err := runner.RunTests(file, testEnvironment, problem.TestCases)
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

func init() {
	testCmd.Flags().StringVarP(&testEnvironment, "environment", "e", "nodejs", "Runtime environment (nodejs)")
}
