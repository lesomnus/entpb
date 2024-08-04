package pbgen

import (
	"fmt"

	"github.com/lesomnus/entpb/pbgen/ident"
)

func OptionGoPackage(v string) Option {
	return Option{
		Name:  ident.Must("go_package"),
		Value: &UnsafeLiteral{Value: fmt.Sprintf(`"%s"`, v)},
	}
}

var (
	FeatureEnumTypeClosed = Option{Name: ident.Must("features", "enum_type"), Value: &UnsafeLiteral{Value: "CLOSED"}}
	FeatureEnumTypeOpen   = Option{Name: ident.Must("features", "enum_type"), Value: &UnsafeLiteral{Value: "OPEN"}}

	FeatureFieldPresenceLegacyRequired = Option{Name: ident.Must("features", "field_presence"), Value: &UnsafeLiteral{Value: "LEGACY_REQUIRED"}}
	FeatureFieldPresenceExplicit       = Option{Name: ident.Must("features", "field_presence"), Value: &UnsafeLiteral{Value: "EXPLICIT"}}
	FeatureFieldPresenceImplicit       = Option{Name: ident.Must("features", "field_presence"), Value: &UnsafeLiteral{Value: "IMPLICIT"}}
)
