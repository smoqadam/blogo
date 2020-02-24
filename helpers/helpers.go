package helpers

import "strings"

// ToSlug convert a string to slug
func ToSlug(s string) string {
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Replace(s, "_", "", -1)
	s = strings.ToLower(s)
	return s
}
