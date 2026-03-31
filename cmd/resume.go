package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"boj/internal/workspace"

	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:   "resume <problem_id>",
	Short: "Resume coding a problem by selecting an existing solution file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		files, err := workspace.FindFiles(id)
		if err != nil {
			return err
		}

		if len(files) == 0 {
			return fmt.Errorf("no solution files found for problem %s\nHint: use `boj code %s -e nodejs` to create one", id, id)
		}

		if len(files) == 1 {
			fmt.Println(files[0])
			return nil
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
			return fmt.Errorf("invalid selection")
		}

		fmt.Println(files[n-1])
		return nil
	},
}
