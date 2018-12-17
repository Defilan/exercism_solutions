package hamming

import "errors"

//Distance is used for this example
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("the inputs must be the same length")
	}

	var count int

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
