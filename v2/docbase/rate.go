package docbase

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Rate represents the rate limit for the current client.
type Rate struct {
	// The number of requests per hour the client is currently limited to.
	Limit int64

	// The number of remaining requests the client can make this hour.
	Remaining int64

	// The time at which the current rate limit will reset.
	Reset Timestamp
}

func (r Rate) String() string {
	return fmt.Sprintf("{Limit:%d,Remaining:%d,Reset:{%s}}", r.Limit, r.Remaining, r.Reset)
}

// formatRateReset formats d to look like "[rate reset in 2s]" or
// "[rate reset in 87m02s]" for the positive durations. And like "[rate limit was reset 87m02s ago]"
// for the negative cases.
func formatRateReset(d time.Duration) string {
	isNegative := d < 0
	if isNegative {
		d *= -1
	}
	secondsTotal := int64(0.5 + d.Seconds())
	minutes := secondsTotal / 60
	seconds := secondsTotal - minutes*60

	var timeString string
	if minutes > 0 {
		timeString = fmt.Sprintf("%dm%02ds", minutes, seconds)
	} else {
		timeString = fmt.Sprintf("%ds", seconds)
	}

	if isNegative {
		return fmt.Sprintf("[rate limit was reset %v ago]", timeString)
	}
	return fmt.Sprintf("[rate reset in %v]", timeString)
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.ParseInt(remaining, 10, 64)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = Timestamp{time.Unix(v, 0)}
		}
	}
	return rate
}
