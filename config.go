package group

import "time"

const (

	// DefaultGroupLimit default query limit
	DefaultGroupLimit = 10

	// MaxGroupLimit max query limit
	MaxGroupLimit = 20

	// MinGrpNameLen the minimum length for a group name
	MinGrpNameLen = 6
)

var (
	// TimeoutRequest a request
	TimeoutRequest = time.Second * 5
)
