package ident

import (
	"strings"
)

type Ident string

type Full []Ident

func (i Full) StringSlice() []string {
	vs := make([]string, len(i))
	for j, v := range i {
		vs[j] = string(v)
	}

	return vs
}

func (i Full) String() string {
	return strings.Join(i.StringSlice(), ".")
}

func Must(vs []string) Full {
	is := make([]Ident, len(vs))
	for i, v := range vs {
		is[i] = Ident(v)
	}

	return is
}
