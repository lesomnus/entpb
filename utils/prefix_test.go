package utils_test

import (
	"testing"

	"github.com/lesomnus/entpb/utils"
	"github.com/stretchr/testify/require"
)

func TestGuessPrefix(t *testing.T) {
	require := require.New(t)

	require.Equal("B_B_", utils.GuessPrefix(
		[]string{"B_B_A", "B_B_B", "B_B_C"},
		[]string{"A", "B", "C"},
	))

	require.Equal("B_B_", utils.GuessPrefix(
		[]string{"B_B_A", "B_B_B", "B_B_C"},
		[]string{"B", "C"},
	))
}
