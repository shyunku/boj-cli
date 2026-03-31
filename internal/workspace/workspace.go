package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"boj/internal/boj"
)

var envExtension = map[string]string{
	"nodejs": "js",
	"node":   "js",
	"python": "py",
	"py":     "py",
	"cpp":    "cpp",
	"c++":    "cpp",
}

type Env struct {
	Name string
	Ext  string
}

// Environments returns the canonical list of supported environments.
func Environments() []Env {
	return []Env{
		{Name: "nodejs", Ext: ".js"},
		{Name: "python", Ext: ".py"},
		{Name: "cpp", Ext: ".cpp"},
	}
}

// EnvFromFile infers the environment from a file's extension.
func EnvFromFile(file string) string {
	switch {
	case strings.HasSuffix(file, ".js"):
		return "nodejs"
	case strings.HasSuffix(file, ".py"):
		return "python"
	case strings.HasSuffix(file, ".cpp"):
		return "cpp"
	default:
		return "nodejs"
	}
}

func ExtForEnv(env string) (string, error) {
	ext, ok := envExtension[strings.ToLower(env)]
	if !ok {
		return "", fmt.Errorf("unsupported environment: %q", env)
	}
	return ext, nil
}

// CreateFile creates a solution file with the problem as a comment header.
func CreateFile(p *boj.Problem, env string) (string, error) {
	ext, err := ExtForEnv(env)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s.%s", p.ID, ext)
	if _, err := os.Stat(filename); err == nil {
		abs, _ := filepath.Abs(filename)
		return abs, nil
	}

	content := buildFile(p, ext)
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}

	abs, _ := filepath.Abs(filename)
	return abs, nil
}

// FindFiles returns all solution files for a given problem ID in the current directory.
func FindFiles(id string) ([]string, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var matches []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasPrefix(e.Name(), id+".") {
			abs, _ := filepath.Abs(e.Name())
			matches = append(matches, abs)
		}
	}
	return matches, nil
}

func buildFile(p *boj.Problem, ext string) string {
	cp := commentPrefix(ext)
	header := buildHeader(p, cp)

	switch ext {
	case "js":
		return header + jsTemplate()
	case "py":
		return header + pyTemplate()
	case "cpp":
		return header + cppTemplate()
	default:
		return header
	}
}

func commentPrefix(ext string) string {
	if ext == "py" {
		return "#"
	}
	return "//"
}

func buildHeader(p *boj.Problem, cp string) string {
	sep := cp + " " + strings.Repeat("=", 60)
	var b strings.Builder

	line := func(s string) {
		if s == "" {
			b.WriteString(cp + "\n")
		} else {
			b.WriteString(cp + " " + s + "\n")
		}
	}

	b.WriteString(sep + "\n")
	line(fmt.Sprintf("Problem %s. %s", p.ID, p.Title))
	b.WriteString(sep + "\n")
	b.WriteString(cp + "\n")

	if p.Description != "" {
		line("[Problem]")
		for _, l := range splitLines(p.Description) {
			line(l)
		}
		b.WriteString(cp + "\n")
	}

	if p.Input != "" {
		line("[Input]")
		for _, l := range splitLines(p.Input) {
			line(l)
		}
		b.WriteString(cp + "\n")
	}

	if p.Output != "" {
		line("[Output]")
		for _, l := range splitLines(p.Output) {
			line(l)
		}
		b.WriteString(cp + "\n")
	}

	for i, tc := range p.TestCases {
		line(fmt.Sprintf("[Test Case %d]", i+1))
		line("Input:")
		for _, l := range splitLines(tc.Input) {
			line("  " + l)
		}
		line("Output:")
		for _, l := range splitLines(tc.Output) {
			line("  " + l)
		}
		b.WriteString(cp + "\n")
	}

	b.WriteString(sep + "\n\n")
	return b.String()
}

func splitLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func jsTemplate() string {
	return `const readline = require('readline');

const rl = readline.createInterface({ input: process.stdin });
const lines = [];
rl.on('line', line => lines.push(line.trim()));
rl.on('close', () => {
  answer(lines);
});

function answer(lines) {
  // TODO: solve problem
  // use console.log() to print your answer
}
`
}

func pyTemplate() string {
	return `import sys
input = sys.stdin.readline

def answer():
    # TODO: solve problem
    pass

answer()
`
}

func cppTemplate() string {
	return `#include <bits/stdc++.h>
using namespace std;

void answer() {
    // TODO: solve problem
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    answer();
    return 0;
}
`
}
