package provider

import "errors"

var (
	// ErrorMustJoinGroup when an user whant to see group and he does not joinned it yet
	ErrorMustJoinGroup = errors.New("You need to join group")

	// ErrorGroupNotExists when the group not exists
	ErrorGroupNotExists = errors.New("The group does not exists")
)
