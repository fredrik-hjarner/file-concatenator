package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCorrectLength(t *testing.T) {
	slice := []string{"1", "2"}
	result := chunk(slice, 1)
	if len(result) != 2 {
		t.Error()
	}
}

func TestCorrectContent(t *testing.T) {
	slice := []string{"1", "2"}
	result := chunk(slice, 1)
	expected := [][]string{[]string{"1"}, []string{"2"}}
	require.EqualValues(t, expected, result)
}
