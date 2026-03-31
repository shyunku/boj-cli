package solvedac

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const perPage = 15

type SearchResponse struct {
	Count int       `json:"count"`
	Items []Problem `json:"items"`
}

type Problem struct {
	ProblemID         int     `json:"problemId"`
	TitleKo           string  `json:"titleKo"`
	Level             int     `json:"level"`
	AcceptedUserCount int     `json:"acceptedUserCount"`
	AverageTries      float64 `json:"averageTries"`
	Tags              []Tag   `json:"tags"`
}

type Tag struct {
	Key          string        `json:"key"`
	DisplayNames []DisplayName `json:"displayNames"`
}

type DisplayName struct {
	Language string `json:"language"`
	Name     string `json:"name"`
	Short    string `json:"short"`
}

func (t Tag) DisplayKo() string {
	for _, d := range t.DisplayNames {
		if d.Language == "ko" {
			return d.Short
		}
	}
	for _, d := range t.DisplayNames {
		if d.Language == "en" {
			return d.Short
		}
	}
	return t.Key
}

func Search(keyword, tier string, page int) (*SearchResponse, error) {
	query := keyword
	if tier != "" {
		query += " tier:" + tier
	}

	apiURL := fmt.Sprintf(
		"https://solved.ac/api/v3/search/problem?query=%s&page=%d&sort=id&direction=asc",
		url.QueryEscape(query), page,
	)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; boj-cli/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to search: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("search failed (status %d)", resp.StatusCode)
	}

	var result SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// TotalPages returns the total number of pages for the given total count.
func TotalPages(count int) int {
	if count == 0 {
		return 1
	}
	return (count + perPage - 1) / perPage
}
