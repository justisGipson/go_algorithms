// extended.go
// description: Implementation of Extended GCD Algorithm
// details:
// A simple implementation of Extended GCD algorithm, that returns GCD, a and b
// which solves ax + by = gcd(a, b)
// author(s) [justis](https://github.com/justisGipson)
// see extended_test.go

package gcd

func Extended(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, xPrime, yPrime := Extended(b%a, a)
	return gcd, yPrime - (b/a)*xPrime, xPrime
}
