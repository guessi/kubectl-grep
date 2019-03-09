package cmd

import (
	"strings"
)

func trimQuoteAndSpace(input string) string {
	if len(input) >= 2 {
		if input[0] == '"' && input[len(input)-1] == '"' {
			return input[1 : len(input)-1]
		}
		if input[0] == '\'' && input[len(input)-1] == '\'' {
			return input[1 : len(input)-1]
		}
	}
	return strings.TrimSpace(input)
}
