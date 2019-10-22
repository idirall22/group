package provider

import "errors"

var (
	// ErrorMustJoinGroup when an user whant to see group and he does not joinned it yet
	ErrorMustJoinGroup = errors.New("You need to join group")

	// ErrorGroupNotExists when the group not exists
	ErrorGroupNotExists = errors.New("The group does not exists")

	// ErrorGroupExists when the group already exists
	ErrorGroupExists = errors.New("Group with the same name already exists")

	// ErrorUpdateGroup when a user want to update a group and the group not exists or
	// the user is not the admin
	ErrorUpdateGroup = errors.New("You can not update the group")

	// ErrorDeleteGroup when a user want to delete a group and the group not exists or
	// the user is not the admin or the group contains members
	ErrorDeleteGroup = errors.New("You can not delete the group")
)
