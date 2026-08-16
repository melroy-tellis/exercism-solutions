package gigasecond

import "time"

// AddGigasecond returns a time.Time value corresponding to the instance occurring 1 gigasecond
// after the passed time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000000000000)
}
