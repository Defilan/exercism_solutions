package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	h = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
 	h['T'] = 0
	for i := 0; i<len(d); i++ {
		switch string(d[i]) {
		case "A":
			h['A']++
		case "C":
			h['C']++
		case "G":
			h['G']++
		case "T":
			h['T']++
		default:
			return nil, errors.New("invalid nucleotides")
		}
	}
	return h, nil
}
