package sqlhelper

import (
	"database/sql"
)

// StrToNullStr converts [string] to [sql.NullString].
// If 's' is empty and 'emptyIsNULL' is true, result's Valid is false.
//
// [string]: https://go.dev/ref/spec#String_types
func StrToNullStr(s string, emptyIsNULL bool) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  !(s == "" && emptyIsNULL),
	}
}

// NullStrToStr converts [sql.NullString] to [string].
// If 'ns.Valid' is false, empty [string] is returned.
//
// [string]: https://go.dev/ref/spec#String_types
func NullStrToStr(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
