package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateAverage(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := (9.0 + 3.5 + 9.0 + 10.0 + 6.0) / 5.0
		result := calculateAverage(9.0, 3.5, 9.0, 10.0, 6.0)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := (5.1 + 4.3 + 10.0 + 9.0) / 4
		result := calculateAverage(5.1, 4.3, 10.0, 9.0)

		require.Equal(t, expected, result)
	})

	t.Run("3", func(t *testing.T) {
		expected := 10.00
		result := calculateAverage(10.0, 10.0, 10.0)

		require.Equal(t, expected, result)
	})
}
