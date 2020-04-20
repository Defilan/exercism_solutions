package raindrops

import (
	"math"
	"strconv"
	"strings"
)
// Convert is the main function of Raindrops, returning the requested information regarding factors.
func Convert(num int) string {
	returner := Newreturn(float64(num))
	if returner == "" {
		return strconv.Itoa(num)
	}
	return returner
}
// Newreturn processes the incoming number and determines if it is a factor.
func Newreturn(num float64) string{
	var forerunner strings.Builder
	numList := []float64{3,5,7}
	wordList := []string{"Pling","Plang","Plong"}
	for i, factor := range numList{
		if num/factor == math.Trunc(num/factor) {
			forerunner.WriteString(wordList[i])
		}
	}
	return forerunner.String()
}