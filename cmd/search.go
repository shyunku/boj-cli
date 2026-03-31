package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"boj/internal/solvedac"

	"github.com/spf13/cobra"
)

var (
	searchTier string
	searchPage int
)

var searchCmd = &cobra.Command{
	Use:   "search <keyword>",
	Short: "Search problems from solved.ac",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword := args[0]
		page := searchPage

		for {
			result, err := solvedac.Search(keyword, searchTier, page)
			if err != nil {
				return err
			}

			if result.Count == 0 {
				fmt.Println("No problems found.")
				return nil
			}

			totalPages := solvedac.TotalPages(result.Count)
			if len(result.Items) > 15 {
				result.Items = result.Items[:15]
			}
			printResults(result)
			fmt.Printf("\n%s\n", strings.Repeat("─", 72))
			fmt.Printf("Page %d / %d  (%d results)\n", page, totalPages, result.Count)

			if totalPages == 1 {
				break
			}

			fmt.Printf("[n]ext  [p]rev  [number] jump  [q]uit > ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch input {
			case "q", "quit", "":
				return nil
			case "n", "next":
				if page < totalPages {
					page++
				} else {
					fmt.Println("already on last page")
				}
			case "p", "prev":
				if page > 1 {
					page--
				} else {
					fmt.Println("already on first page")
				}
			default:
				n, err := strconv.Atoi(input)
				if err != nil || n < 1 || n > totalPages {
					fmt.Printf("invalid page (1-%d)\n", totalPages)
				} else {
					page = n
				}
			}
		}

		return nil
	},
}

func printResults(result *solvedac.SearchResponse) {
	fmt.Println()
	for _, p := range result.Items {
		tier := solvedac.TierInfo(p.Level)

		badge := fmt.Sprintf("[%s]", tier.Label)
		coloredBadge := tier.Color + padRight(badge, 4) + tier.Reset

		id := fmt.Sprintf("%d.", p.ProblemID)
		title := truncate(p.TitleKo, 28)
		idTitle := padRight(id+" "+title, 36)

		tags := buildTags(p.Tags, 30)
		ac := formatAC(p.AcceptedUserCount)

		fmt.Printf("%s %-36s  %-30s  %8s AC\n", coloredBadge, idTitle, tags, ac)
	}
}

func buildTags(tags []solvedac.Tag, maxLen int) string {
	if len(tags) == 0 {
		return ""
	}
	var names []string
	for _, t := range tags {
		names = append(names, t.DisplayKo())
	}
	joined := "[" + strings.Join(names, ", ") + "]"
	return truncate(joined, maxLen)
}

func formatAC(n int) string {
	if n >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	}
	if n >= 1000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return strconv.Itoa(n)
}

// truncate cuts s to maxLen runes, appending "…" if cut.
func truncate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) <= maxLen {
		return s
	}
	runes := []rune(s)
	return string(runes[:maxLen-1]) + "…"
}

// padRight pads s with spaces to width w (rune-aware).
func padRight(s string, w int) string {
	n := utf8.RuneCountInString(s)
	if n >= w {
		return s
	}
	return s + strings.Repeat(" ", w-n)
}

func init() {
	searchCmd.Flags().StringVarP(&searchTier, "tier", "t", "", "Filter by tier (e.g. bronze, silver, gold, gold3)")
	searchCmd.Flags().IntVarP(&searchPage, "page", "p", 1, "Starting page")
}
