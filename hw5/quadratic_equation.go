package hw5

import (
	"errors"
	"math"
)

var (
	ErrNoRoots              = errors.New("there are no real roots")
	ErrDiscriminantLessZero = errors.New("the discriminant is negative, then there are no real roots")
)

const delta = 0.000001

func Equation(a, b, c float64) (x1, x2 float64, err error) {
	if math.Abs(a) <= delta {
		return 0, 0, ErrNoRoots
	}

	D := b*b - 4*a*c
	if D < -delta {
		return 0, 0, ErrDiscriminantLessZero
	}

	if math.Abs(D) <= delta {
		x1 = -b / 2 * a
		return x1, x1, nil
	}
	// D > delta
	sqrt := math.Sqrt(D)
	return (-b + sqrt) / (2 * a), (-b - sqrt) / (2 * a), nil
}
