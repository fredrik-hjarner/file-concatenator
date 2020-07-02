package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCorrectLength(t *testing.T) {
	slice := []string{"1", "2"}
	result := Chunk(slice, 1)
	if len(result) != 2 {
		t.Error()
	}
}

func TestCorrectContent(t *testing.T) {
	slice := []string{"1", "2"}
	result := Chunk(slice, 1)
	expected := [][]string{[]string{"1"}, []string{"2"}}
	require.EqualValues(t, expected, result)
}

func TestCorrectContent2(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5"}
	result := Chunk(slice, 2)
	expected := [][]string{
		[]string{"1", "2"},
		[]string{"3", "4"},
		[]string{"5"},
	}
	require.EqualValues(t, expected, result)
}
