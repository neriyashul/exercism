// The package add a gigasecond to time t
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond return time + 1,000,000,000 seconds
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
