package group

import "time"

const (

	// DefaultGroupOffset default query offset
	DefaultGroupOffset = 10

	// MaxGroupOffset max query offset
	MaxGroupOffset = 20

	// MinGrpNameLen the minimum length for a group name
	MinGrpNameLen = 6
)

var (
	// TimeoutRequest a request
	TimeoutRequest = time.Second * 5
)
