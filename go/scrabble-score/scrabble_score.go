package scrabble

import "strings"

//Score will calculate the Scrabble score of an incoming string and return its score as an integer
func Score(letters string) int {
	points := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}
	returner := 0
	for i := 0; i < len(letters); i++ {
		for letter, value := range points {
			for y := 0; y < len(letter); y++ {
				scrabblechar := string(letters[i])
				if strings.ToLower(scrabblechar) == string(letter[y]) {
					returner += value
				}
			}
		}
	}
	return returner
}
