package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/entpb/pbgen"
	"github.com/stretchr/testify/require"
)

func TestEnum(t *testing.T) {
	t.Run("single field", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Enum{
			Name: "Role",
			Body: []pbgen.EnumBody{
				pbgen.EnumField{Name: "UNSPECIFIED", Number: 0},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`enum Role {
	UNSPECIFIED = 0;
}`, v)
	})

	t.Run("multiple fields", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Enum{
			Name: "Role",
			Body: []pbgen.EnumBody{
				pbgen.EnumField{Name: "UNSPECIFIED", Number: 0},
				pbgen.EnumField{Name: "OWNER", Number: 1},
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`enum Role {
	UNSPECIFIED = 0;
	OWNER = 1;
}`, v)
	})
}
