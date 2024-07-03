package entpb

import "entgo.io/ent/schema"

const KeyAnnotationLabel = "Key"

func Key(key string, num int) schema.Annotation {
	return &KeyAnnotation{Key: key, Number: num}
}

type KeyAnnotation struct {
	Key    string
	Number int
}

func (KeyAnnotation) Name() string {
	return KeyAnnotationLabel
}
