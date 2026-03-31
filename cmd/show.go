package cmd

import (
	"fmt"

	"boj/internal/boj"

	"github.com/spf13/cobra"
)

var showTestcases bool

var showCmd = &cobra.Command{
	Use:   "show <problem_id>",
	Short: "Show problem details (fetches and caches from BOJ)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		problem, err := boj.GetProblem(id)
		if err != nil {
			return err
		}

		fmt.Printf("# %s. %s\n\n", problem.ID, problem.Title)
		fmt.Printf("## Problem\n%s\n\n", problem.Description)
		fmt.Printf("## Input\n%s\n\n", problem.Input)
		fmt.Printf("## Output\n%s\n\n", problem.Output)

		if showTestcases {
			if len(problem.TestCases) == 0 {
				fmt.Println("No test cases found.")
				return nil
			}
			for i, tc := range problem.TestCases {
				fmt.Printf("## Test Case %d\n", i+1)
				fmt.Printf("Input:\n%s\n\n", tc.Input)
				fmt.Printf("Output:\n%s\n\n", tc.Output)
			}
		}

		return nil
	},
}

func init() {
	showCmd.Flags().BoolVarP(&showTestcases, "testcases", "t", false, "Show example test cases")
}
