// package caesar the shift cipher
// ref: https://en.wikipedia.org/wiki/Caesar_cipher

package caesar

func Encrypt(input string, key int) string {
	// if key is neg val
	// updates key, the number which matches to key modulo 26
	var key8 uint8 = byte(key%26+26) % 26

	outputBuffer := []byte{}
	// r is a rune equivalent of uint32
	for _, r := range input {
		newbyte := byte(r)
		if 'A' <= r && r <= 'Z' {
			outputBuffer = append(outputBuffer, 'A'+(newbyte-'A'+key8)%26)
		} else if 'a' <= r && r <= 'z' {
			outputBuffer = append(outputBuffer, 'a'+(newbyte-'a'+key8)%26)
		} else {
			outputBuffer = append(outputBuffer, newbyte)
		}
	}
	return string(outputBuffer)
}

// decrypts by left shift of each key, each char of input
func Decrypt(input string, key int) string {
	// left shift of key is the same as right shift of 26 key
	return Encrypt(input, 26-key)
}
