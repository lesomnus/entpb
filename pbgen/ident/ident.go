package ident

import (
	"fmt"
	"strings"
)

type Ident string

func (i Ident) Full() Full {
	return Full{Segments: []string{string(i)}}
}

type Full struct {
	Segments []string
	Braced   bool
}

func (i Full) String() string {
	v := strings.Join(i.Segments, ".")
	if i.Braced {
		v = fmt.Sprintf("(%s)", v)
	}

	return v
}

func Must(v string, vs ...string) Full {
	return Full{
		Segments: append([]string{v}, vs...),
	}
}

func (i Full) WithBraces() Full {
	i.Braced = true
	return i
}

func (i Full) Last() Ident {
	if len(i.Segments) == 0 {
		return ""
	}

	return Ident(i.Segments[len(i.Segments)-1])
}

func (i Full) Append(vs ...Ident) Full {
	for _, v := range vs {
		i.Segments = append(i.Segments, string(v))
	}

	return i
}

func (i Full) Equals(other Full) bool {
	return i.String() == other.String()
}
