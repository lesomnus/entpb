package entpb

import (
	"fmt"

	"entgo.io/ent/schema"
	"github.com/mitchellh/mapstructure"
)

func decodeAnnotation[T schema.Annotation](v T, annotations map[string]any) (T, bool) {
	a, ok := annotations[v.Name()]
	if !ok {
		return v, false
	}
	if err := mapstructure.Decode(a, v); err != nil {
		panic(fmt.Errorf("decode %s: %w", v.Name(), err))
	}

	return v, true
}
