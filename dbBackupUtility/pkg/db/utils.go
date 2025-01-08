package database

import "strings"

func escapeSQL(s string) string {
	return strings.Replace(s, "'", "''", -1)
}