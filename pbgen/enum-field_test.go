package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/entpb/pbgen"
	"github.com/stretchr/testify/require"
)

func TestEnumField(t *testing.T) {
	t.Run("normal field", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.EnumField{
			Name:   "UNSPECIFIED",
			Number: 0,
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal("UNSPECIFIED = 0;", v)
	})
}
