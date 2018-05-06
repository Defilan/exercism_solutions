// Package gigasecond is a simple app that takes a date, adds a gigasecond, and returns a date
package gigasecond

import "time"

// AddGigasecond takes in a date, adds a gigasecond to it, and returns a date
func AddGigasecond(t time.Time) time.Time {
	returner := t.Add(time.Second * (1000000000))
	return returner
}
