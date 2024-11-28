package stats

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumFunction(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 2.5
		result := minimumFunction(10.0, 4.5, 9.8, 7.6, 7.6, 2.5)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 4.0
		result := minimumFunction(9.0, 4.0, 6.5, 4.5)

		require.Equal(t, expected, result)
	})
}

func Test_maximumFunction(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 10.0
		result := maximumFunction(10.0, 4.5, 9.8, 7.6, 7.6, 2.5)
		
		require.Equal(t, expected, result)
	})
	
	t.Run("2", func(t *testing.T) {
		expected := 9.0
		result := maximumFunction(9.0, 4.0, 6.5, 4.5)

		require.Equal(t, expected, result)
	})
}

func Test_averageFunction(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 6.0
		result := averageFunction(9.0, 4.0, 6.5, 4.5)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 7.0
		result := averageFunction(10.0, 4.5, 9.8, 7.6, 7.6, 2.5)

		require.Equal(t, expected, result)
	})
}