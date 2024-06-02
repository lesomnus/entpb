package pbgen

func OptionGoPackage(v string) Option {
	return Option{
		Name:  []Ident{"go_package"},
		Value: v,
	}
}

var (
	FeatureFieldPresenceLegacyRequired = Option{Name: []Ident{"features", "field_presence"}, Value: "LEGACY_REQUIRED"}
	FeatureFieldPresenceExplicit       = Option{Name: []Ident{"features", "field_presence"}, Value: "EXPLICIT"}
	FeatureFieldPresenceImplicit       = Option{Name: []Ident{"features", "field_presence"}, Value: "IMPLICIT"}
)
