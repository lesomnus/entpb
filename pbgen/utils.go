package pbgen

import "strings"

func ParseFullIndent(v string) FullIdent {
	return strings.Split(v, ".")
}
