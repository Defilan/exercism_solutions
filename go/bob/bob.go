// Package bob returns responses by Bob, the cranky teen
package bob

import "strings"

// Hey is the main function of Bob, returning Bob's whitty retorts.
func Hey(remark string) string {

	greeting := strings.TrimSpace(remark)

	switch {
	case len(greeting) == 0:
		return "Fine. Be that way!"
	case greeting == strings.ToUpper(greeting) && greeting != strings.ToLower(greeting) && strings.Contains(greeting, "?"):
		return "Calm down, I know what I'm doing!"
	case greeting == strings.ToUpper(greeting) && greeting != strings.ToLower(greeting):
		return "Whoa, chill out!"
	case strings.HasSuffix(greeting, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
