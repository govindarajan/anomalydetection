package throttle

import (
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
)

// Throttler defines interface for throttling
type Throttler interface {
	Throttle(types.Context) bool
}

// basicThrottler type for mocking Throttler
type basicThrottler struct {
	Unit  string
	Limit int
}

// BasicThrottler creates basic throtlle
func BasicThrottler(limit int, unit string) Throttler {
	return &basicThrottler{Unit: unit, Limit: limit}
}

// Throttle implementing throttle interface
func (bt basicThrottler) Throttle(ctx types.Context) bool {
	// if time.Now().Unix()%7 == 0 {
	// 	return false
	// }
	return true
}
