// package caesar the shift cipher
// ref: https://en.wikipedia.org/wiki/Caesar_cipher

package caesar

func Encrypt(input string, key int) string {
	// if key is neg val
	// updates key, the number which matches to key modulo 26
	key = (key%26 + 26) % 26

	outputBuffer = []byte{}

	for _, r := range input {
		newbyte := byte(r)
		if 'A' <= newbyte && newbyte <= 'Z' {
			outputBuffer = append(outputBuffer, byte(int('A')+int(int(newbyte-'A')+key)%26))
		} else if 'a' <= newbyte && newbyte <= 'z' {
			outputBuffer = append(outputBuffer, byte(int('a')+int(int(newbyte-'a')+key)%26))
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
