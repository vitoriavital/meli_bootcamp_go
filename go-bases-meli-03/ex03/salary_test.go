package salary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateSalary(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 9000.0
		result := calculateSalary(120, "A")

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 5400.0
		result := calculateSalary(180, "B")

		require.Equal(t, expected, result)
	})

	t.Run("3", func(t *testing.T) {
		expected := 2000.0
		result := calculateSalary(120, "C")

		require.Equal(t, expected, result)
	})
}
