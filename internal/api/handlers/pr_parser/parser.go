package pr_parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid PR title format")

type PrTitle struct {
	IssueId   string
	IssueType string
	AppLayer  string
	Title     string
	Version   string
}

func ParsePRTitle(prTitle string) (*PrTitle, error) {
	fmt.Println("Parsing PR Title:", prTitle)

	pattern := `^\(([^)]+)\) \| ([A-Z]+)\[([A-Za-z]+)\]: (.+) \| ([\d.]+)$`
	re := regexp.MustCompile(pattern)

	parts := strings.Split(prTitle, "|")
	issueId := strings.ReplaceAll(parts[0], "(", "")
	issueId = strings.ReplaceAll(issueId, ")", "")

	fmt.Println("Extracted Issue ID:", issueId)
	matches := re.FindStringSubmatch(prTitle)
	fmt.Println("Matches found:", matches, len(matches))
	if len(matches) != 6 {
		return nil, ErrInvalidFormat
	}

	return &PrTitle{
		IssueId:   matches[1],
		IssueType: matches[2],
		AppLayer:  matches[3],
		Title:     matches[4],
		Version:   matches[5],
	}, nil
}
