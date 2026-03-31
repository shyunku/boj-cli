package runner

import (
	"bytes"
	"fmt"
	"os"
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
	case "python", "py":
		return runWithPython(file, testCases)
	case "cpp", "c++":
		return runWithCpp(file, testCases)
	default:
		return nil, fmt.Errorf("unsupported environment: %q (see `boj langs`)", env)
	}
}

func runEach(command string, args []string, testCases []boj.TestCase) ([]Result, error) {
	results := make([]Result, len(testCases))
	for i, tc := range testCases {
		cmd := exec.Command(command, args...)
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
			results[i] = Result{Passed: false, Expected: expected, Actual: errMsg}
			continue
		}

		actual := strings.TrimSpace(stdout.String())
		results[i] = Result{Passed: actual == expected, Expected: expected, Actual: actual}
	}
	return results, nil
}

func runWithNode(file string, testCases []boj.TestCase) ([]Result, error) {
	return runEach("node", []string{file}, testCases)
}

func runWithPython(file string, testCases []boj.TestCase) ([]Result, error) {
	python := "python3"
	if _, err := exec.LookPath(python); err != nil {
		python = "python"
	}
	return runEach(python, []string{file}, testCases)
}

func runWithCpp(file string, testCases []boj.TestCase) ([]Result, error) {
	// Compile to a temp binary
	tmp, err := os.CreateTemp("", "boj-cpp-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	tmp.Close()
	defer os.Remove(tmp.Name())

	compile := exec.Command("g++", "-O2", "-o", tmp.Name(), file)
	var compileErr bytes.Buffer
	compile.Stderr = &compileErr
	if err := compile.Run(); err != nil {
		return nil, fmt.Errorf("compilation failed:\n%s", compileErr.String())
	}

	return runEach(tmp.Name(), nil, testCases)
}
