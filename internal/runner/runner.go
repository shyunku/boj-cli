package runner

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"boj/internal/boj"
)

type Result struct {
	Passed   bool
	Expected string
	Actual   string
}

func RunTests(file, env string, testCases []boj.TestCase) ([]Result, error) {
	switch env {
	case "nodejs", "node":
		return runWithNode(file, testCases)
	default:
		return nil, fmt.Errorf("unsupported environment: %q (supported: nodejs)", env)
	}
}

func runWithNode(file string, testCases []boj.TestCase) ([]Result, error) {
	results := make([]Result, len(testCases))

	for i, tc := range testCases {
		cmd := exec.Command("node", file)
		cmd.Stdin = strings.NewReader(tc.Input)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		expected := strings.TrimSpace(tc.Output)

		if err := cmd.Run(); err != nil {
			errMsg := strings.TrimSpace(stderr.String())
			if errMsg == "" {
				errMsg = err.Error()
			}
			results[i] = Result{
				Passed:   false,
				Expected: expected,
				Actual:   errMsg,
			}
			continue
		}

		actual := strings.TrimSpace(stdout.String())
		results[i] = Result{
			Passed:   actual == expected,
			Expected: expected,
			Actual:   actual,
		}
	}

	return results, nil
}
