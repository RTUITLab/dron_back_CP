package role

import "errors"

var (
	ErrRoleNotValid = errors.New("role not valid")
)

type RoleEnum string

const (
	Student 	= "student"
	Teacher 	= "teacher"
	Admin		= "admin"
)

func ValidRole(role string) error {
	switch role {
	case Student, Teacher, Admin:
		return nil
	default:
		return ErrRoleNotValid
	}
}