package strand

import (
	"strings"
)

//ToRNA takes a DNA strand and returns its RNA complement
func ToRNA(dna string) string {
	var rna strings.Builder
	for x := 0; x< len(dna); x++ {
		switch string(dna[x]) {
		case "G":
			rna.WriteString("C")
		case "C":
			rna.WriteString("G")
		case "T":
			rna.WriteString("A")
		case "A":
			rna.WriteString("U")
		}
	}
	return rna.String()
}
