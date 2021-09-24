// Package xor is an encryption algorithm that operates the exclusive disjunction(XOR)
// ref: https://en.wikipedia.org/wiki/XOR_cipher

package xor

// Encrypt encrypts with Xor encryption after converting each character to byte
// The returned value might not be readable because there is no guarantee
// which is within the ASCII range
// If using other type such as string, []int, or some other types,
// add the statements for converting the type to []byte.

func Encrypt(key byte, plaintext []byte) []byte {
	cipherText := []byte{}
	for _, ch := range plaintext {
		cipherText = append(cipherText, key^ch)
	}
	return cipherText
}

func Decrypt(key byte, ciphertext []byte) []byte {
	plaintext := []byte{}
	for _, ch := range ciphertext {
		plaintext = append(plaintext, key^ch)
	}
	return plaintext
}
