// rsa.go
// simple rsa algorithm implementation
// details:
// A simple RSA Encryption and Decryption algorithm.
// It uses prime numbers that fit in int64 datatypes and
// thus both the Encrypt and Decrypt are not a production
// ready implementation. The OpenSSL implementation of RSA
// also adds a padding which is not present in this algorithm.
// author(s) [Taj](https://github.com/tjgurwara99)
// see rsa_test.go

package rsa

import (
	"errors"

	modular "github.com/TheAlgorithms/Go/math/modular"
)

// raised when Encrypt func fails on message
var ErrorFailedEncrypt = errors.New("failed to encrypt")

// raised when Decrypt func fails on message
var ErrorFailedDecrypt = errors.New("failed to decrypt")

// rsa alo encryption - uses modular exponentiation in math dir
func Encrypt(message []rune, pubExponent, modulus int64) ([]rune, error) {
	var encrypted []rune

	for _, letter := range message {
		encryptedLetter, err := modular.Exponentiation(int64(letter), pubExponent, modulus)
		if err != nil {
			return nil, ErrorFailedEncrypt
		}
		encrypted = append(encrypted, rune(encryptedLetter))
	}
	return encrypted, nil
}

// decrypts encrypted rune slice based on RSA algorithm
func Decrypt(encrypted []rune, privExponent, modulus int64) (string, error) {
	var decrypted []rune

	for _, letter := range encrypted {
		decryptedLetter, err := modular.Exponentiation(int64(letter), privExponent, modulus)
		if err != nil {
			return nil, ErrorFailedDecrypt
		}
		decrypted = append(decrypted, rune(decryptedLetter))
	}
	return string(decrypted), nil
}
