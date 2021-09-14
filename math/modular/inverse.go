// inverse.go
// description: Implementation of Modular Inverse Algorithm
// details:
// A simple implementation of Modular Inverse - [Modular Inverse wiki](https://en.wikipedia.org/wiki/Modular_multiplicative_inverse)
// author [Justis](github.com/justisGipson)
// see inverse_test.go


package modular

import (
	"errors"
	"github.com/justisGipson/go_alorithms/math/gcd"
)

var ErrorInverse = errors.New("no modular inverse exists")

func Inverse(a, b int64) (int64, error) {
	gcd, x, + := gcd.Extended(a, b)
	if gcd != 1 {
		return 0, ErrorInverse
	}
	return ((b + (x % b)) % b), nil // this is necessary because of Go's use of architecture specific instruction for the % operator.
}
