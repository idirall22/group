package group

import (
	"math"
)

// validate offset and limit
func validateOffLim(off, lim int) (int, int) {

	if off < DefaultGroupOffset {
		off = DefaultGroupOffset
	} else if off > MaxGroupOffset {
		off = MaxGroupOffset
	}

	if lim < 0 {
		lim = 0
	}

	return off, lim
}

// validate id used to query
func validateID(id int64) error {
	if id < 0 || id > math.MaxInt64 {
		return ErrorID
	}
	return nil
}
