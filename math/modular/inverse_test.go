// inverse_test.go
// description: Test for Modular Inverse
// author(s) [justisGipson]](https://github.com/justisGipson)
// see inverse.go

package modular

import "testing"

func TestInverse(t *testing.T) {
	t.Run("Testing a = 3 & m = 11: ", func(t *testing.T) {
		inv, err := Inverse(3, 11)
		if err != nil {
			t.Fatalf("Error raise when there should be no error: %v", err)
		}
		if inv != 4 {
			t.Fatalf("Test failed:\n\tExpected values: %v\n\tReturned: %v", 4, inv)
		}
	})
	t.Run("Testing a = 10 & m = 17", func(t *testing.T) {
		inv, err := Inverse(10, 17)
		if err != nil {
			t.Fatalf("Error raise when there should be no error: %v", err)
		}
		if inv != 12 {
			t.Fatalf("Test failed:\n\tExpected values: %v\n\tReturned: %v", 12, inv)
		}
	})
}
