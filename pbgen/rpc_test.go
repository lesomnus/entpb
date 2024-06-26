package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/entpb/pbgen"
	"github.com/lesomnus/entpb/pbgen/ident"
	"github.com/stretchr/testify/require"
)

func TestRpc(t *testing.T) {
	t.Run("unary", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: ident.Full{"User"}},
			Response: pbgen.RpcType{Type: ident.Full{"User"}},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (User);`, v)
	})

	t.Run("stream request", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: ident.Full{"User"}, Stream: true},
			Response: pbgen.RpcType{Type: ident.Full{"User"}},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (stream User) returns (User);`, v)
	})

	t.Run("stream response", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.Rpc{
			Name:     "Create",
			Request:  pbgen.RpcType{Type: ident.Full{"User"}},
			Response: pbgen.RpcType{Type: ident.Full{"User"}, Stream: true},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`rpc Create (User) returns (stream User);`, v)
	})
}
