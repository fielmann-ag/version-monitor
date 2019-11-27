package testing

import (
	"time"
)

// Time is a static test time
var Time time.Time

func init() {
	var err error
	Time, err = time.Parse(time.RFC822, "12 Dec 25 15:00 UTC")
	if err != nil {
		panic(err)
	}
}
