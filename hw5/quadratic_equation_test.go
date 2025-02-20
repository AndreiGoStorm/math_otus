package hw5

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

var a, b, c float64

func TestAEquilZero(t *testing.T) {
	t.Run("than A equal zero", func(t *testing.T) {
		// Arrange
		a = 0
		b = rand.Float64()
		c = rand.Float64()

		// Act
		_, _, err := Equation(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrNoRoots)
	})

	t.Run("than A equal delta", func(t *testing.T) {
		// Arrange
		a = delta
		b = rand.Float64()
		c = rand.Float64()

		// Act
		_, _, err := Equation(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrNoRoots)
	})

	t.Run("than a less delta", func(t *testing.T) {
		// Arrange
		a = delta / 2
		b = rand.Float64()
		c = rand.Float64()

		// Act
		_, _, err := Equation(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrNoRoots)
	})
}

func TestDiscriminantLessZero(t *testing.T) {
	t.Run("discriminant less then zero", func(t *testing.T) {
		// Arrange
		a = 1.25
		b = -5.4351
		c = 6.001

		// Act
		_, _, err := Equation(a, b, c)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrDiscriminantLessZero)
	})
}

func TestDiscriminantEqualZero(t *testing.T) {
	t.Run("discriminant equal zero with delta", func(t *testing.T) {
		// Arrange
		a = 2
		b = 4
		c = 2.00000001
		expectedX := -4

		// Act
		x1, x2, err := Equation(a, b, c)

		// Assert
		require.NoError(t, err)
		require.InDelta(t, expectedX, x1, delta)
		require.InDelta(t, expectedX, x2, delta)
	})
}

func TestDiscriminantMoreZero(t *testing.T) {
	t.Run("discriminant more then zero with delta", func(t *testing.T) {
		// Arrange
		a = 4
		b = -16
		c = 10
		expectedX1 := 3.224744
		expectedX2 := 0.775255

		// Act
		x1, x2, err := Equation(a, b, c)

		// Assert
		require.NoError(t, err)
		require.InDelta(t, expectedX1, x1, delta)
		require.InDelta(t, expectedX2, x2, delta)
	})

	t.Run("discriminant more then zero with decreased delta", func(t *testing.T) {
		// Arrange
		a = 2
		b = 7
		c = -11
		expectedX1 := 1.176
		expectedX2 := -4.6761

		// Act
		x1, x2, err := Equation(a, b, c)

		// Assert
		require.NoError(t, err)
		require.InDelta(t, expectedX1, x1, delta*1000)
		require.InDelta(t, expectedX2, x2, delta*100)
	})
}
