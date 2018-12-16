package hamming

import "errors"

//Distance is used for this example
func Distance(a, b string) (int, error) {
	count := 0

	if len(a) != len(b) {
		return -1, errors.New("The inputs must be the same length")
	}

	for i, c := range a {
		if (string(c)) != (string(b[i])) {
			count++
		}
	}
	return count, nil
}
