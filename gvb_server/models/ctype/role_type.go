package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin        = 1
	PermissionUser         = 2
	PermissionVisitor      = 3
	PermissionDisabledUser = 4 // Changed DisableUser to DisabledUser for clarity
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() any {
	var str string
	switch s {
	case PermissionAdmin:
		str = "Admin"
	case PermissionUser:
		str = "User"
	case PermissionVisitor:
		str = "Visitor"
	case PermissionDisabledUser:
		str = "Muted"
	default:
		str = "others"
	}
	return str
}
