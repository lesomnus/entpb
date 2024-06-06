package utils_test

import (
	"testing"

	"github.com/lesomnus/entpb/utils"
	"github.com/stretchr/testify/require"
)

func TestSnake(t *testing.T) {
	require := require.New(t)

	require.Equal("a_b", utils.Snake("a_b"))
	require.Equal("a_bc", utils.Snake("a_bc"))
	require.Equal("a_b_c", utils.Snake("a_b_c"))

	require.Equal("a_b", utils.Snake("AB"))
	require.Equal("a_b_c", utils.Snake("ABC"))
	require.Equal("a_bc", utils.Snake("ABc"))
	require.Equal("ab_c", utils.Snake("AbC"))
	require.Equal("abc", utils.Snake("Abc"))
	require.Equal("a_b_c", utils.Snake("aBC"))
	require.Equal("a_bc", utils.Snake("aBc"))
	require.Equal("ab_c", utils.Snake("abC"))
	require.Equal("abc", utils.Snake("abc"))
}

func TestTitle(t *testing.T) {
	require := require.New(t)

	require.Equal("AB", utils.Pascal("a_b"))
	require.Equal("ABC", utils.Pascal("a_b_C"))
	require.Equal("AbC", utils.Pascal("Ab_c"))
}
