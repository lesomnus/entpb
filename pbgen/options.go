package pbgen

import (
	"fmt"

	"github.com/lesomnus/entpb/pbgen/ident"
)

func OptionGoPackage(v string) Option {
	return Option{
		Name:  ident.Full{"go_package"},
		Value: fmt.Sprintf(`"%s"`, v),
	}
}

var (
	FeatureEnumTypeClosed = Option{Name: ident.Full{"features", "enum_type"}, Value: "CLOSED"}
	FeatureEnumTypeOpen   = Option{Name: ident.Full{"features", "enum_type"}, Value: "OPEN"}

	FeatureFieldPresenceLegacyRequired = Option{Name: ident.Full{"features", "field_presence"}, Value: "LEGACY_REQUIRED"}
	FeatureFieldPresenceExplicit       = Option{Name: ident.Full{"features", "field_presence"}, Value: "EXPLICIT"}
	FeatureFieldPresenceImplicit       = Option{Name: ident.Full{"features", "field_presence"}, Value: "IMPLICIT"}
)
