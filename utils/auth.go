package utils

import "strings"

func ConvertRolesToString(roles []string) string {
	return strings.Join(roles, ",")
}

func ConvertStringToRoles(roles string) []string {
	return strings.Split(roles, ",")
}
