package taxes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculateSalaryTaxes(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 0.00
		result := calculateSalaryTaxes(40000.00)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 60000.00 * 0.17
		result := calculateSalaryTaxes(60000.00)

		require.Equal(t, expected, result)
	})

	t.Run("3", func(t *testing.T) {
		expected := 160000.00 * 0.27
		result := calculateSalaryTaxes(160000.00)

		require.Equal(t, expected, result)
	})
}
