package pr_parser

import (
	"errors"
	"regexp"
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
	pattern := `^\(([^)]+)\) \| ([A-Z]+)\[([A-Za-z]+)\]: (.+) \| ([\d.]+)$`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(prTitle)
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
