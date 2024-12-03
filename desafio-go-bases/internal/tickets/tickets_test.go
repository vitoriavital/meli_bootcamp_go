package tickets_test

import (
	"testing"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
	"errors"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 14
		result,_ := tickets.GetTotalTickets("Japan")

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 11
		result, _ := tickets.GetTotalTickets("Mexico")

		require.Equal(t, expected, result)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 151
		result,_ := tickets.GetCountByPeriod("night")

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 256
		result, _ := tickets.GetCountByPeriod("morning")

		require.Equal(t, expected, result)
	})
	t.Run("3", func(t *testing.T) {
		expected := 304
		result,_ := tickets.GetCountByPeriod("early morning")

		require.Equal(t, expected, result)
	})

	t.Run("4", func(t *testing.T) {
		expected := 289
		result, _ := tickets.GetCountByPeriod("afternoon")

		require.Equal(t, expected, result)
	})
	t.Run("5", func(t *testing.T) {
		expected := errors.New("Error: invalid is not a valid period of time!")
		_, err := tickets.GetCountByPeriod("invalid")

		require.Equal(t, expected, err)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expected := 1.40
		result,_ := tickets.AverageDestination("Japan", 1000)

		require.Equal(t, expected, result)
	})

	t.Run("2", func(t *testing.T) {
		expected := 1.10
		result, _ := tickets.AverageDestination("Mexico", 1000)

		require.Equal(t, expected, result)
	})
	t.Run("3", func(t *testing.T) {
		expected := 4.50
		result,_ := tickets.AverageDestination("Brazil", 1000)

		require.Equal(t, expected, result)
	})

	t.Run("4", func(t *testing.T) {
		expected := 0.0
		result, _ := tickets.AverageDestination("not real destination", 1000)

		require.Equal(t, expected, result)
	})
}
