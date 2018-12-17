package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	if rhyme == nil {
		return []string{}
	}

	var builder []string
	count := len(rhyme)
	for i := range rhyme {
		var s string
		if i != count-1 {
			s = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		} else {
			s = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		}
		builder = append(builder, s)
	}
	return builder
}
