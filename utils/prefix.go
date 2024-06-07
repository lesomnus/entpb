package utils

import (
	"slices"
	"strings"
)

func GuessPrefix(prefixed []string, values []string) string {
	for _, v := range values {
		for _, u := range prefixed {
			prefix, ok := strings.CutSuffix(u, v)
			if !ok {
				continue
			}

			n := 0
			for _, w := range prefixed {
				if slices.Contains(values, strings.TrimPrefix(w, prefix)) {
					n++
				}
			}
			if n == len(values) {
				return prefix
			}
		}
	}

	return ""
}
