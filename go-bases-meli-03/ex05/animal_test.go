package animal

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func Test_dogFoodAmount(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 200.0
		result := dogFoodAmount(20)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 400.0
		result := dogFoodAmount(40)

		require.Equal(t, expected, result)
	})
}

func Test_catFoodAmount(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 20.0
		result := catFoodAmount(4)
		
		require.Equal(t, expected, result)
	})
	
	t.Run("2", func(t *testing.T) {
		expected := 35.0
		result := catFoodAmount(7)

		require.Equal(t, expected, result)
	})
}

func Test_hamsterFoodAmount(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 1.0
		result := hamsterFoodAmount(4)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 2.5
		result := hamsterFoodAmount(10)

		require.Equal(t, expected, result)
	})
}

func Test_spiderFoodAmount(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 1.2
		result := spiderFoodAmount(8)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 12.0
		result := spiderFoodAmount(80)

		require.Equal(t, expected, result)
	})
}