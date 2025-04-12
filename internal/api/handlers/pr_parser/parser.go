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
	issueId = strings.TrimSpace(issueId)

	typeAndLayer := strings.Split(parts[1], ":")
	typeAndLayerSplit := strings.Split(typeAndLayer[0], "[")
	issueTypeValue := strings.TrimSpace(typeAndLayerSplit[0])

	layer := typeAndLayerSplit[1]
	layer = strings.ReplaceAll(layer, "]", "")
	layer = strings.TrimSpace(layer)

	title := typeAndLayer[1]
	title = strings.TrimSpace(title)

	version := strings.TrimSpace(parts[2])

	fmt.Println("Extracted issueType:", issueTypeValue)
	fmt.Println("Extracted Issue ID:", issueId)
	fmt.Println("Extracted Layer:", layer)
	fmt.Println("Extracted Title:", title)
	fmt.Println("Extracted Version:", version)
	matches := re.FindStringSubmatch(prTitle)
	fmt.Println("Matches found:", matches, len(matches))
	// if len(matches) != 6 {
	// 	return nil, ErrInvalidFormat
	// }

	return &PrTitle{
		IssueId:   issueId,
		IssueType: issueTypeValue,
		AppLayer:  layer,
		Title:     title,
		Version:   version,
	}, nil
}
