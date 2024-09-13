package search

import (
	"fmt"
	"regexp"
	"strings"
)

type RuleWithDescription struct {
	Regex       *regexp.Regexp
	Description string
}

func SearchStringInResults(search string, results []string, sensitive bool, path string) {
	for _, str := range results {
		if sensitive {
			// Case-sensitive search
			if strings.Contains(str, search) {
				fmt.Println(str)
			}
		} else {
			// Case-insensitive search
			if strings.Contains(strings.ToLower(str), strings.ToLower(search)) {
				fmt.Printf("%s\n\nFound inside: %s\n", str, path)
			}
		}
	}
}

func SearchRegexInResults(pattern string, results []string, sensitive bool) {
	var re *regexp.Regexp
	var err error

	if sensitive {
		re, err = regexp.Compile(pattern)
	} else {
		// For case-insensitive search, use the (?i) flag in the regex pattern
		re, err = regexp.Compile("(?i)" + pattern)
	}

	if err != nil {
		fmt.Println("Invalid regex pattern:", err)
		return
	}

	for _, str := range results {
		if re.MatchString(str) {
			fmt.Println(str)
		}
	}
}

// SearchWithRegexes applies multiple regex patterns to each result and prints matching strings.
func SearchWithRegexes(rules map[string]RuleWithDescription, results []string, excludeIDs []string, path string) {
	// Create a set of excluded rule IDs for fast lookup
	excludeSet := make(map[string]struct{}, len(excludeIDs))
	for _, id := range excludeIDs {
		excludeSet[id] = struct{}{}
	}

	for _, str := range results {
		for id, rule := range rules {
			// Skip rules that are in the exclude set
			if _, excluded := excludeSet[id]; excluded {
				continue
			}

			if rule.Regex.MatchString(str) {
				fmt.Printf("'%s' (%s): %s\nFound inside: %s\n\n", id, rule.Description, str, path)
			}
		}
	}
}