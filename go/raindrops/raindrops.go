package raindrops

import (
	"strconv"
)

// Convert is the main function of Raindrops, returning the requested information regarding factors.
func Convert(num int) string {
	returner := ""
	if num%3 == 0 {
		returner += "Pling"
	}
	if num%5 == 0 {
		returner += "Plang"
	}
	if num%7 == 0 {
		returner += "Plong"
	}
	if returner == "" {
		returner += strconv.Itoa(num)
	}
	return returner
}
