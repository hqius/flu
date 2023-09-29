package db

import "strings"

func IsDuplicate(err error) bool {
	return strings.Contains(err.Error(), DUPLICATE_ENTRY)
}
